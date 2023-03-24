package actions

type PJSIPShowRegistrationInboundContactStatusesAction struct {
	ActionID string
}

func (PJSIPShowRegistrationInboundContactStatusesAction) ActionTypeName() string {
	return "PJSIPShowRegistrationInboundContactStatuses"
}
func (a PJSIPShowRegistrationInboundContactStatusesAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowRegistrationInboundContactStatusesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPShowRegistrationInboundContactStatuses(opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowRegistrationInboundContactStatusesAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
