package ami

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/parsidev/asterisk/ami/message/actions"
	er "github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const attrActionID = "ActionID"

var errClose = errors.New("Close")

type Conn struct {
	g       *errgroup.Group
	conn    net.Conn
	reader  *bufio.Reader
	nextID  func() string
	pending chan *asyncMsg
	recv    chan *Message
	ctx     context.Context
	closer  context.CancelFunc
	closed  bool

	version string
	logger  *zap.Logger
	conf    *ConnectOptions
	subs    []*subscribe
	subLoc  sync.Mutex
}

func (c *Conn) read(ctx context.Context) (err error) {
	log := c.logger
	log.Debug("start read loop")
	defer func() {
		log.Debug("done read loop")
	}()
	for {
		msg := &Message{}

		if err = msg.Read(c.reader); err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
			return
		}

		if msg.Type != "" {
			c.recv <- msg
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (c *Conn) loop(ctx context.Context) (err error) {
	ids := map[string]*asyncMsg{}
	defer func() {
		for _, v := range ids {
			v.err = err
			v.complete()
		}
	}()

	log := c.logger
	cleanTicker := time.NewTicker(5 * time.Second)

	c.logger.Debug("start event loop")
	defer func() {
		log.Debug("done event loop")
	}()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-cleanTicker.C:
			c.cleanUnsub()
		case async := <-c.pending:
			//err = c.conn.SetDeadline(time.Now().Add(1 * time.Second))
			//if err != nil {
			//	return
			//}

			// send pending message
			msg := async.msg
			//log.Sugar().With("id", async.id, "type", msg.Type, "name", msg.Name).Debug("send message")

			err = msg.Write(c.conn)

			if err != nil {
				async.err = err
				async.complete()
			} else if async.id != "" {
				ids[async.id] = async
			}
			c.onSend(async)
		case msg := <-c.recv:
			// received
			if msg.Type == MessageTypeResponse {
				id := msg.AttrString(attrActionID)
				if id != "" {
					async := ids[id]
					if async == nil {
						log.Sugar().With("id", id).Warn("response untracked")
					} else {
						async.resp = msg
						async.complete()
					}
				}
			}
			c.onRecv(msg)
		}
	}
}

func (c *Conn) onSend(msg *asyncMsg) {
	subs := c.subs

	for _, v := range subs {
		if v.unsub || !v.onSent {
			continue
		}
		ctx := v.ctx
		if ctx == nil {
			ctx = c.ctx
		}
		v.unsub = !v.f(ctx, msg.msg)
	}
}

func (c *Conn) onRecv(msg *Message) {
	subs := c.subs
	for _, v := range subs {
		if v.unsub || !v.onRecv {
			continue
		}
		ctx := v.ctx
		if ctx == nil {
			ctx = c.ctx
		}
		// NOTE blocked
		v.unsub = !v.f(ctx, msg)
	}
}

func (c *Conn) Close() error {
	if c.closer != nil {
		c.closed = true
		c.closer()
		c.closer = nil
		err := c.g.Wait()
		if errors.Is(err, errClose) {
			return nil
		}
		return err
	}
	if c.closed {
		return nil
	}
	return errors.New("not init")
}

func (c *Conn) dial(addr string) (err error) {
	conf := c.conf

	if conf.AllowReconnect {
		// NOTE reconnect keep pending, but fail all async
		go func() {
			log := c.logger
			onErr := conf.OnConnectErr
			if onErr == nil {
				onErr = func(conn *Conn, err error) {
				}
			}

			var err error
			for !c.closed {
				err = c.dialOnce(addr)
				if err != nil {
					log.Sugar().With("err", err).Warn("ami.Conn: dial")

					onErr(c, err)
					// fixme improve wait strategy
					<-time.NewTimer(time.Second).C
					continue
				}
				if conf.OnConnected != nil {
					conf.OnConnected(c)
					log.Sugar().Info("ami.Conn: connected")
				}
				err = c.g.Wait()
				if err != nil {
					log.Sugar().With("err", err).Warn("ami.Conn: error")
					onErr(c, err)
				}
				c.g = nil
			}
			log.Sugar().Info("ami.Conn: stop reconnect, conn closed")
		}()
		return nil
	}
	return c.dialOnce(addr)
}

func (c *Conn) dialOnce(addr string) (err error) {
	conf := c.conf
	conn, err := conf.Dialer.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if e := conn.Close(); e != nil {
				err = multierror.Append(err, e)
			}
		}
	}()
	return c.connect(conn)
}

func (c *Conn) connect(conn net.Conn) (err error) {
	log := c.logger
	r := bufio.NewReader(conn)
	c.reader = r
	c.conn = conn

	line, err := r.ReadString('\n')
	if err != nil {
		return er.Wrap(err, "scan ami initial line")
	}

	// check connection
	match := regexp.MustCompile("Asterisk Call Manager/([0-9.]+)").FindStringSubmatch(line)
	if len(match) > 1 {
		c.version = match[1]
		log.Sugar().With("version", c.version).Debug("AMI Version")
	} else {
		err = errors.New(fmt.Sprintf("Invalid server header: %q", line))
		return
	}

	ctx := c.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	c.g, ctx = errgroup.WithContext(ctx)
	c.g.Go(func() error {
		return c.read(ctx)
	})
	c.g.Go(func() error {
		return c.loop(ctx)
	})
	// manual close
	waitCtx, closer := context.WithCancel(ctx)
	c.g.Go(func() error {
		<-ctx.Done()
		closer()
		return conn.Close()
	})
	c.g.Go(func() error {
		<-waitCtx.Done()
		return errClose
	})
	c.closer = closer

	conf := c.conf
	if conf.Username != "" {
		var resp *Message
		resp, err = c.Request(actions.LoginAction{
			UserName: conf.Username,
			Secret:   conf.Secret,
		}, RequestTimeout(2*time.Second))

		if err != nil {
			err = er.Wrap(err, "request login")
		} else if !resp.Success() {
			err = er.Wrap(resp.Error(), "login")
		}
		if err != nil {
			log.Sugar().With("err", err).Info("login failed")
			return err
		}
		log.Info("login success")
	}
	return
}

func Connect(addr string, opts ...ConnectOption) (conn *Conn, err error) {
	opt := &ConnectOptions{
		Timeout: 10 * time.Second,
		Context: context.Background(),
		Logger:  zap.L(),
	}

	for _, v := range opts {
		if err = v(opt); err != nil {
			return nil, err
		}
	}
	if opt.Dialer == nil {
		opt.Dialer = &net.Dialer{
			Timeout: opt.Timeout,
		}
	}

	var id uint64
	conn = &Conn{
		ctx:     opt.Context,
		conf:    opt,
		logger:  opt.Logger,
		recv:    make(chan *Message, 4096),
		pending: make(chan *asyncMsg, 100),
		nextID: func() string {
			return fmt.Sprint(atomic.AddUint64(&id, 1))
		},
	}
	for _, sub := range opt.subscribers {
		_, err = conn.Subscribe(sub.sub, sub.opts...)
		if err != nil {
			return nil, err
		}
	}
	return conn, conn.dial(addr)
}
