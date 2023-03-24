package actions

import "github.com/parsidev/asterisk/ami/message"

type MonitorAction struct {
	ActionID string
	Channel  string
	File     string
	Format   string
	Mix      string
}

func (MonitorAction) ActionTypeName() string {
	return "Monitor"
}
func (a MonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *MonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Monitor(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MonitorAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
