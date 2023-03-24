package actions

import "github.com/parsidev/asterisk/ami/message"

type QueueSummaryAction struct {
	ActionID string
	Queue    string
}

func (QueueSummaryAction) ActionTypeName() string {
	return "QueueSummary"
}
func (a QueueSummaryAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueSummaryAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueSummary(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueSummaryAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
