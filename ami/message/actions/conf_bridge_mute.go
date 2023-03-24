package actions

import "github.com/parsidev/asterisk/ami/message"

type ConfbridgeMuteAction struct {
	ActionID   string
	Conference string
	Channel    string
}

func (ConfbridgeMuteAction) ActionTypeName() string {
	return "ConfbridgeMute"
}
func (a ConfbridgeMuteAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeMuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeMute(conference string, channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ConfbridgeMuteAction{
		Conference: conference,
		Channel:    channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
