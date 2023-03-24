package actions

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
func (cli *Client) QueueReset(opts ...RequestOption) (res *Response, err error) {
	req := &QueueResetAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
