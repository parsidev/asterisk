package actions

import "github.com/parsidev/asterisk/ami/message"

type QueueRemoveAction struct {
	ActionID  string
	Queue     string
	Interface string
}

func (QueueRemoveAction) ActionTypeName() string {
	return "QueueRemove"
}
func (a QueueRemoveAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueRemoveAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueRemove(queue string, iface string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueRemoveAction{
		Queue:     queue,
		Interface: iface,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
