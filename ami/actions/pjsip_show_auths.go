package actions

type PJSIPShowAuthsAction struct {
	ActionID string
}

func (PJSIPShowAuthsAction) ActionTypeName() string {
	return "PJSIPShowAuths"
}
func (a PJSIPShowAuthsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowAuthsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PJSIPShowAuths(opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowAuthsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
