package actions

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
func (cli *Client) QueueStatus(opts ...RequestOption) (res *Response, err error) {
	req := &QueueStatusAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
