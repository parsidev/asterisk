package actions

type PJSIPShowAorsAction struct {
	ActionID string
}

func (PJSIPShowAorsAction) ActionTypeName() string {
	return "PJSIPShowAors"
}
func (a PJSIPShowAorsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowAorsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PJSIPShowAors(opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowAorsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
