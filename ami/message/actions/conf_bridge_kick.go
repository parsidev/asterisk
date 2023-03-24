package actions

import "github.com/parsidev/asterisk/ami/message"

type ConfbridgeKickAction struct {
	ActionID   string
	Conference string
	Channel    string
}

func (ConfbridgeKickAction) ActionTypeName() string {
	return "ConfbridgeKick"
}
func (a ConfbridgeKickAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeKickAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeKick(conference string, channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ConfbridgeKickAction{
		Conference: conference,
		Channel:    channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
