package actions

import "github.com/parsidev/asterisk/ami/message"

type PauseMonitorAction struct {
	ActionID string
	Channel  string
}

func (PauseMonitorAction) ActionTypeName() string {
	return "PauseMonitor"
}
func (a PauseMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *PauseMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PauseMonitor(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PauseMonitorAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
