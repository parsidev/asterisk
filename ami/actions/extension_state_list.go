package actions

type ExtensionStateListAction struct {
	ActionID string
}

func (ExtensionStateListAction) ActionTypeName() string {
	return "ExtensionStateList"
}
func (a ExtensionStateListAction) GetActionID() string {
	return a.ActionID
}
func (a *ExtensionStateListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) ExtensionStateList(opts ...RequestOption) (res *Response, err error) {
	req := &ExtensionStateListAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
