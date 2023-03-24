package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) ModuleLoad(loadType string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ModuleLoadAction{
		LoadType: loadType,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
