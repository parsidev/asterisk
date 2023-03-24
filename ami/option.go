package ami

import (
	"context"
	"net"
	"time"

	"go.uber.org/zap"
)

type CustomDialer interface {
	Dial(network, address string) (net.Conn, error)
}

type ConnErrHandler func(*Conn, error)
type ConnHandler func(*Conn)

type ConnectOptions struct {
	Context        context.Context
	Timeout        time.Duration
	AllowReconnect bool
	Username       string // login username
	Secret         string // login secret
	Logger         *zap.Logger
	Dialer         CustomDialer
	OnConnectErr   ConnErrHandler
	OnConnected    ConnHandler
	subscribers    []struct {
		sub  SubscribeFunc
		opts []SubscribeOption
	}
}
type ConnectOption func(c *ConnectOptions) error

func WithAuth(username string, secret string) ConnectOption {
	return func(c *ConnectOptions) error {
		c.Username = username
		c.Secret = secret
		return nil
	}
}

func WithSubscribe(cb SubscribeFunc, opts ...SubscribeOption) ConnectOption {
	return func(c *ConnectOptions) error {
		c.subscribers = append(c.subscribers, struct {
			sub  SubscribeFunc
			opts []SubscribeOption
		}{sub: cb, opts: opts})
		return nil
	}
}
