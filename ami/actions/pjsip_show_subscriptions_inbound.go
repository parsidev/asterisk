package actions

import "github.com/parsidev/asterisk/ami/message"

type PJSIPShowSubscriptionsInboundAction struct {
	ActionID string
}

func (PJSIPShowSubscriptionsInboundAction) ActionTypeName() string {
	return "PJSIPShowSubscriptionsInbound"
}
func (a PJSIPShowSubscriptionsInboundAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowSubscriptionsInboundAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPShowSubscriptionsInbound(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowSubscriptionsInboundAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
