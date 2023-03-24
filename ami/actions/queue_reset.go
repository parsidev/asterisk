package actions

import "github.com/parsidev/asterisk/ami/message"

type QueueResetAction struct {
	ActionID string
	Queue    string
}

func (QueueResetAction) ActionTypeName() string {
	return "QueueReset"
}
func (a QueueResetAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueResetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueReset(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueResetAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
