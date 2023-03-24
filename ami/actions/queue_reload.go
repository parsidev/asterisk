package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) QueueReload(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueReloadAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
