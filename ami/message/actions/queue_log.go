package actions

import "github.com/parsidev/asterisk/ami/message"

type QueueLogAction struct {
	ActionID  string
	Queue     string
	Event     string
	Uniqueid  string
	Interface string
	Message   string
}

func (QueueLogAction) ActionTypeName() string {
	return "QueueLog"
}
func (a QueueLogAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueLogAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueLog(queue string, event string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueLogAction{
		Queue: queue,
		Event: event,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
