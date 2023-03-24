package actions

import "github.com/parsidev/asterisk/ami/message"

type MixMonitorAction struct {
	ActionID string
	Channel  string
	File     string
	Options  string
	Command  string
}

func (MixMonitorAction) ActionTypeName() string {
	return "MixMonitor"
}
func (a MixMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *MixMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MixMonitor(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MixMonitorAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
