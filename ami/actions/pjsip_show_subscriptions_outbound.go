package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) PJSIPShowSubscriptionsOutbound(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowSubscriptionsOutboundAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
