package actions

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
func (cli *Client) QueueSummary(opts ...RequestOption) (res *Response, err error) {
	req := &QueueSummaryAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
