package actions

import "github.com/parsidev/asterisk/ami/message"

type ChangeMonitorAction struct {
	ActionID string
	Channel  string
	File     string
}

func (ChangeMonitorAction) ActionTypeName() string {
	return "ChangeMonitor"
}
func (a ChangeMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *ChangeMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ChangeMonitor(channel string, file string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ChangeMonitorAction{
		Channel: channel,
		File:    file,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
