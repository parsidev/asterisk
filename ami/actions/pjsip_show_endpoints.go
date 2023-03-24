package actions

type PJSIPShowEndpointsAction struct {
	ActionID string
}

func (PJSIPShowEndpointsAction) ActionTypeName() string {
	return "PJSIPShowEndpoints"
}
func (a PJSIPShowEndpointsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowEndpointsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PJSIPShowEndpoints(opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowEndpointsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
