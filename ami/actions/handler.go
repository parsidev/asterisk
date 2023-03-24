package actions

import "github.com/parsidev/asterisk/ami/message"

type Handler interface {
	Request(req *message.Request) *message.Response
}
type HandlerFunc func(req *message.Request) *message.Response

func (f HandlerFunc) Request(req *message.Request) *message.Response {
	return f(req)
}

type Client struct {
	Handler
}

func (cli *Client) Action(act Action, res interface{}, opts ...message.RequestOption) error {
	r, err := message.RequestBuilder(act, res, opts...)
	if err != nil {
		return err
	}
	return cli.Request(r).Error
}
