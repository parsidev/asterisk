package actions

import "github.com/parsidev/asterisk/ami/message"

type MixMonitorMuteAction struct {
	ActionID  string
	Channel   string
	Direction string
	State     string
}

func (MixMonitorMuteAction) ActionTypeName() string {
	return "MixMonitorMute"
}
func (a MixMonitorMuteAction) GetActionID() string {
	return a.ActionID
}
func (a *MixMonitorMuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MixMonitorMute(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MixMonitorMuteAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
