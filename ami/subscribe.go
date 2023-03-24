package ami

import (
	"context"
)

type subscribe struct {
	f      SubscribeFunc
	ctx    context.Context
	unsub  bool
	onSent bool
	onRecv bool
}
type SubscribeFunc func(ctx context.Context, message *Message) bool

type SubscribeOption func(o *subscribe) error

func SubscribeSend() SubscribeOption {
	return func(o *subscribe) error {
		o.onSent = true
		return nil
	}
}

func SubscribeSetContext(ctx context.Context) SubscribeOption {
	return func(o *subscribe) error {
		o.ctx = ctx
		return nil
	}
}

func (c *Conn) Subscribe(cb SubscribeFunc, opts ...SubscribeOption) (func(), error) {
	c.subLoc.Lock()
	defer c.subLoc.Unlock()

	sub := &subscribe{
		f:      cb,
		onRecv: true,
	}
	for _, v := range opts {
		err := v(sub)
		if err != nil {
			return nil, err
		}
	}

	c.subs = append(c.subs, sub)

	return func() {
		sub.unsub = true
	}, nil
}
func (c *Conn) cleanUnsub() {
	c.subLoc.Lock()
	defer c.subLoc.Unlock()
	var neo []*subscribe
	subs := c.subs
	for _, sub := range subs {
		if !sub.unsub {
			neo = append(neo, sub)
		}
	}
	c.subs = neo
}

func SubscribeChan(c chan *Message, names ...string) SubscribeFunc {
	m := make(map[string]struct{}, len(names))
	for _, v := range names {
		m[v] = struct{}{}
	}
	filter := func(ctx context.Context, msg *Message) bool {
		_, ok := m[msg.Name]
		return ok
	}
	if len(names) == 0 {
		filter = func(ctx context.Context, msg *Message) bool {
			return true
		}
	}
	return func(ctx context.Context, msg *Message) bool {
		if filter(ctx, msg) {
			c <- msg
		}

		return true
	}
}

func SubscribeFullyBootedChanOnce(c chan *Message) SubscribeFunc {
	return func(ctx context.Context, msg *Message) bool {
		if msg.Name == "FullyBooted" {
			c <- msg
			return false
		}
		return true
	}
}
