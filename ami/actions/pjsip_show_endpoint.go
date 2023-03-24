package actions

import "github.com/parsidev/asterisk/ami/message"

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

func (cli *Client) PJSIPShowEndpoint(endPoint string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowEndpointAction{
		Endpoint: endPoint,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
