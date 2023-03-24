package actions

import "github.com/parsidev/asterisk/ami/message"

type PRIDebugFileSetAction struct {
	ActionID string
	File     string
}

func (PRIDebugFileSetAction) ActionTypeName() string {
	return "PRIDebugFileSet"
}
func (a PRIDebugFileSetAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIDebugFileSetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PRIDebugFileSet(file string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PRIDebugFileSetAction{
		File: file,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
