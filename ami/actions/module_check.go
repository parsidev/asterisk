package actions

type ModuleCheckAction struct {
	ActionID string
	Module   string
}

func (ModuleCheckAction) ActionTypeName() string {
	return "ModuleCheck"
}
func (a ModuleCheckAction) GetActionID() string {
	return a.ActionID
}
func (a *ModuleCheckAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ModuleCheck(module string, opts ...RequestOption) (res *Response, err error) {
	req := &ModuleCheckAction{
		Module: module,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
