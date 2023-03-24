package actions

import "github.com/parsidev/asterisk/ami/message"

type UnpauseMonitorAction struct {
	ActionID string
	Channel  string
}

func (UnpauseMonitorAction) ActionTypeName() string {
	return "UnpauseMonitor"
}
func (a UnpauseMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *UnpauseMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) UnpauseMonitor(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &UnpauseMonitorAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
