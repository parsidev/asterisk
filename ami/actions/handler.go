package actions

type Handler interface {
	Request(req *Request) *Response
}
type HandlerFunc func(req *Request) *Response

func (f HandlerFunc) Request(req *Request) *Response {
	return f(req)
}

type Client struct {
	Handler
}

func (cli *Client) Action(act Action, res interface{}, opts ...RequestOption) error {
	r, err := RequestBuilder(act, res, opts...)
	if err != nil {
		return err
	}
	return cli.Request(r).Error
}
