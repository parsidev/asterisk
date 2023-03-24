package actions

import "github.com/parsidev/asterisk/ami/message"

type DAHDIHangupAction struct {
	ActionID     string
	DAHDIChannel string
}

func (DAHDIHangupAction) ActionTypeName() string {
	return "DAHDIHangup"
}
func (a DAHDIHangupAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIHangupAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIHangup(dAHDIChannel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DAHDIHangupAction{
		DAHDIChannel: dAHDIChannel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
