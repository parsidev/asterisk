package actions

import "github.com/parsidev/asterisk/ami/message"

type QueueStatusAction struct {
	ActionID string
	Queue    string
	Member   string
}

func (QueueStatusAction) ActionTypeName() string {
	return "QueueStatus"
}
func (a QueueStatusAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueStatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueStatus(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueStatusAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
