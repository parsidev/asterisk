package actions

import "github.com/parsidev/asterisk/ami/message"

type ReloadAction struct {
	ActionID string
	Module   string
}

func (ReloadAction) ActionTypeName() string {
	return "Reload"
}
func (a ReloadAction) GetActionID() string {
	return a.ActionID
}
func (a *ReloadAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Reload(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ReloadAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
