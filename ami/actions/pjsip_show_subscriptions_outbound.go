package actions

type PJSIPShowSubscriptionsOutboundAction struct {
	ActionID string
}

func (PJSIPShowSubscriptionsOutboundAction) ActionTypeName() string {
	return "PJSIPShowSubscriptionsOutbound"
}
func (a PJSIPShowSubscriptionsOutboundAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowSubscriptionsOutboundAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPShowSubscriptionsOutbound(opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowSubscriptionsOutboundAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
