package actions

type PJSIPShowRegistrationsInboundAction struct {
	ActionID string
}

func (PJSIPShowRegistrationsInboundAction) ActionTypeName() string {
	return "PJSIPShowRegistrationsInbound"
}
func (a PJSIPShowRegistrationsInboundAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowRegistrationsInboundAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPShowRegistrationsInbound(opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowRegistrationsInboundAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
