package actions

import "github.com/parsidev/asterisk/ami/message"

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

func (cli *Client) PJSIPShowEndpoints(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowEndpointsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
