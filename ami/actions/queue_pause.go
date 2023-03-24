package actions

import "github.com/parsidev/asterisk/ami/message"

type QueuePauseAction struct {
	ActionID  string
	Interface string
	Paused    string
	Queue     string
	Reason    string
}

func (QueuePauseAction) ActionTypeName() string {
	return "QueuePause"
}
func (a QueuePauseAction) GetActionID() string {
	return a.ActionID
}
func (a *QueuePauseAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueuePause(iface string, paused string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueuePauseAction{
		Interface: iface,
		Paused:    paused,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
