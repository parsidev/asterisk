package actions

import "github.com/parsidev/asterisk/ami/message"

type StopMixMonitorAction struct {
	ActionID     string
	Channel      string
	MixMonitorID string
}

func (StopMixMonitorAction) ActionTypeName() string {
	return "StopMixMonitor"
}
func (a StopMixMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *StopMixMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) StopMixMonitor(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &StopMixMonitorAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
