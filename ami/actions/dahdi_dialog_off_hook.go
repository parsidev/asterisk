package actions

import "github.com/parsidev/asterisk/ami/message"

type DAHDIDialOffhookAction struct {
	ActionID     string
	DAHDIChannel string
	Number       string
}

func (DAHDIDialOffhookAction) ActionTypeName() string {
	return "DAHDIDialOffhook"
}
func (a DAHDIDialOffhookAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDialOffhookAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIDialOffhook(dAHDIChannel string, number string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DAHDIDialOffhookAction{
		DAHDIChannel: dAHDIChannel,
		Number:       number,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
