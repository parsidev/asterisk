package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) ModuleCheck(module string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ModuleCheckAction{
		Module: module,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
