package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) PJSIPShowRegistrationsInbound(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowRegistrationsInboundAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
