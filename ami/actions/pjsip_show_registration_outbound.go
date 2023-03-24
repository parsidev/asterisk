package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) PJSIPShowRegistrationsOutbound(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowRegistrationsOutboundAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
