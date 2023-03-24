package actions

type ModuleLoadAction struct {
	ActionID string
	Module   string
	LoadType string
}

func (ModuleLoadAction) ActionTypeName() string {
	return "ModuleLoad"
}
func (a ModuleLoadAction) GetActionID() string {
	return a.ActionID
}
func (a *ModuleLoadAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ModuleLoad(loadType string, opts ...RequestOption) (res *Response, err error) {
	req := &ModuleLoadAction{
		LoadType: loadType,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
