package actions

type PJSIPShowRegistrationsOutboundAction struct {
	ActionID string
}

func (PJSIPShowRegistrationsOutboundAction) ActionTypeName() string {
	return "PJSIPShowRegistrationsOutbound"
}
func (a PJSIPShowRegistrationsOutboundAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowRegistrationsOutboundAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPShowRegistrationsOutbound(opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowRegistrationsOutboundAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
