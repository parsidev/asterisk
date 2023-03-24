package actions

type QueueReloadAction struct {
	ActionID   string
	Queue      string
	Members    string
	Rules      string
	Parameters string
}

func (QueueReloadAction) ActionTypeName() string {
	return "QueueReload"
}
func (a QueueReloadAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueReloadAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueReload(opts ...RequestOption) (res *Response, err error) {
	req := &QueueReloadAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
