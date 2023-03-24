package actions

import "github.com/parsidev/asterisk/ami/message"

type HangupAction struct {
	ActionID string
	Channel  string
	Cause    string
}

func (HangupAction) ActionTypeName() string {
	return "Hangup"
}
func (a HangupAction) GetActionID() string {
	return a.ActionID
}
func (a *HangupAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Hangup(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &HangupAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
