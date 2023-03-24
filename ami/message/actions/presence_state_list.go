package actions

import "github.com/parsidev/asterisk/ami/message"

type PresenceStateListAction struct {
	ActionID string
}

func (PresenceStateListAction) ActionTypeName() string {
	return "PresenceStateList"
}
func (a PresenceStateListAction) GetActionID() string {
	return a.ActionID
}
func (a *PresenceStateListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PresenceStateList(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PresenceStateListAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
