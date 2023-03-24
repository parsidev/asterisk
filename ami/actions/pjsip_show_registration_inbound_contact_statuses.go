package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) PJSIPShowRegistrationInboundContactStatuses(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowRegistrationInboundContactStatusesAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
