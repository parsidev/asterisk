package actions

import "github.com/parsidev/asterisk/ami/message"

type QueueAddAction struct {
	ActionID       string
	Queue          string
	Interface      string
	Penalty        string
	Paused         string
	MemberName     string
	StateInterface string
}

func (QueueAddAction) ActionTypeName() string {
	return "QueueAdd"
}
func (a QueueAddAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueAddAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueAdd(queue string, iface string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueAddAction{
		Queue:     queue,
		Interface: iface,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
