package agi

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	conf       *Options
	logger     *zap.Logger
	listener   net.Listener
	cons       map[*conn]struct{}
	mu         sync.Mutex
	inShutdown bool
	handlers   map[string]Route
}

type Route func(Request)

func NewServer(addr string, opts ...Option) (srv *Server, err error) {

	opt := &Options{
		Addr:         addr,
		Logger:       zap.L(),
		MaxReadBytes: 10 * 1024,
	}

	for _, v := range opts {
		if err = v(opt); err != nil {
			return nil, err
		}
	}

	srv = &Server{
		conf:   opt,
		logger: opt.Logger,
	}

	return
}

func (srv *Server) AddRoute(name string, route Route) error {
	if srv.handlers == nil {
		srv.handlers = make(map[string]Route)
	}

	_, isExist := srv.handlers[name]
	if isExist {
		return errors.New("this route already exist")
	}
	srv.handlers[name] = route
	return nil
}

func (srv *Server) ListenAndServe() {
	addr := srv.conf.Addr
	if addr == "" {
		addr = ":4030"
	}
	srv.logger.Info(fmt.Sprintf("starting server on %v\n", addr))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	defer listener.Close()
	srv.listener = listener
	for {

		if srv.inShutdown {
			break
		}

		if newConn, err := listener.Accept(); err == nil {
			conn := &conn{
				Conn:          newConn,
				MaxReadBuffer: srv.conf.MaxReadBytes,
			}
			srv.trackConn(conn)
			go srv.handle(conn)
		}
	}
	return
}

func (srv *Server) trackConn(c *conn) {
	defer srv.mu.Unlock()
	srv.mu.Lock()
	if srv.cons == nil {
		srv.cons = make(map[*conn]struct{})
	}
	srv.cons[c] = struct{}{}
}

func (srv *Server) handle(conn *conn) {
	defer func() {
		_ = conn.Close()
		srv.deleteConn(conn)
	}()
	conn.reader = bufio.NewReader(conn)
	conn.writer = bufio.NewWriter(conn)
	scanner := bufio.NewScanner(conn.reader)
	var str strings.Builder

	sc := make(chan bool)
	for {
		go func(s chan bool) {
			s <- scanner.Scan()
		}(sc)
		select {
		case scanned := <-sc:
			if !scanned {
				if err := scanner.Err(); err != nil {
					return
				}
				return
			}
			if scanner.Text() == "" {
				request := Parse(str.String(), conn, srv)

				handler, isExist := srv.handlers[request.networkString]
				if !isExist {
					return
				}
				handler(request)
				_ = conn.Close()
				srv.deleteConn(conn)
			} else {
				str.WriteString("\n")
				str.WriteString(scanner.Text())
			}
		}
	}
}

func (srv *Server) deleteConn(conn *conn) {
	defer srv.mu.Unlock()
	srv.mu.Lock()
	delete(srv.cons, conn)
}

func (srv *Server) Shutdown() {
	srv.inShutdown = true
	srv.logger.Info("shutting down...")
	_ = srv.listener.Close()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			srv.logger.Info(fmt.Sprintf("waiting on %v connections", len(srv.cons)))
		}
		if len(srv.cons) == 0 {
			return
		}
	}
}

func Parse(text string, conn *conn, srv *Server) Request {
	request := Request{
		conn: conn,
		srv:  srv,
		args: map[string]string{},
	}
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		data := strings.Split(line, ": ")
		if len(data) < 2 {
			continue
		}
		switch strings.TrimSpace(data[0]) {
		case "agi_network":
			request.network = strings.TrimSpace(data[1])
		case "agi_network_script":
			request.networkString = strings.TrimSpace(data[1])
		case "agi_request":
			request.request = strings.TrimSpace(data[1])
		case "agi_channel":
			request.channel = strings.TrimSpace(data[1])
		case "agi_language":
			request.language = strings.TrimSpace(data[1])
		case "agi_type":
			request.aType = strings.TrimSpace(data[1])
		case "agi_uniqueid":
			request.uniqueID = strings.TrimSpace(data[1])
		case "agi_version":
			request.version = strings.TrimSpace(data[1])
		case "agi_callerid":
			request.callerID = strings.TrimSpace(data[1])
		case "agi_calleridname":
			request.callerIdName = strings.TrimSpace(data[1])
		case "agi_callingpres":
			request.calleringPres = strings.TrimSpace(data[1])
		case "agi_callingani2":
			request.calleringAni2 = strings.TrimSpace(data[1])
		case "agi_callington":
			request.calleringTon = strings.TrimSpace(data[1])
		case "agi_callingtns":
			request.calleringTns = strings.TrimSpace(data[1])
		case "agi_dnid":
			request.dnID = strings.TrimSpace(data[1])
		case "agi_rdnis":
			request.rdnis = strings.TrimSpace(data[1])
		case "agi_context":
			request.context = strings.TrimSpace(data[1])
		case "agi_extension":
			request.extension = strings.TrimSpace(data[1])
		case "agi_priority":
			request.priority = strings.TrimSpace(data[1])
		case "agi_enhanced":
			request.enhanced = strings.TrimSpace(data[1])
		case "agi_accountcode":
			request.accountCode = strings.TrimSpace(data[1])
		case "agi_threadid":
			request.threadID = strings.TrimSpace(data[1])
		default:
			request.args[strings.TrimSpace(data[0])] = strings.TrimSpace(data[1])
		}
	}
	return request
}
