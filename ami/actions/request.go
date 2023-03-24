package actions

type Request struct {
	Action   Action
	Response *Response
}

type RequestOption func(r *Request) error

func RequestBuilder(act Action, res interface{}, opts ...RequestOption) (r *Request, err error) {
	r = &Request{
		Action: act,
	}
	for _, v := range opts {
		if err = v(r); err != nil {
			return nil, err
		}
	}
	return
}
