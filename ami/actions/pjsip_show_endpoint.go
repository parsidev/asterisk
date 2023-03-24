package actions

type PJSIPShowEndpointAction struct {
	ActionID string
	Endpoint string
}

func (PJSIPShowEndpointAction) ActionTypeName() string {
	return "PJSIPShowEndpoint"
}
func (a PJSIPShowEndpointAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowEndpointAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PJSIPShowEndpoint(endPoint string, opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowEndpointAction{
		Endpoint: endPoint,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
