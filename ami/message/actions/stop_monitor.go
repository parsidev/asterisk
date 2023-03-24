package actions

import "github.com/parsidev/asterisk/ami/message"

type StopMonitorAction struct {
	ActionID string
	Channel  string
}

func (StopMonitorAction) ActionTypeName() string {
	return "StopMonitor"
}
func (a StopMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *StopMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) StopMonitor(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &StopMonitorAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
