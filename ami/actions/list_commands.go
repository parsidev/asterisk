package actions

type ListCommandsAction struct {
	ActionID string
}

func (ListCommandsAction) ActionTypeName() string {
	return "ListCommands"
}
func (a ListCommandsAction) GetActionID() string {
	return a.ActionID
}
func (a *ListCommandsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ListCommands(opts ...RequestOption) (res *Response, err error) {
	req := &ListCommandsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
