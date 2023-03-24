package actions

import "github.com/parsidev/asterisk/ami/message"

type ConfbridgeUnlockAction struct {
	ActionID   string
	Conference string
}

func (ConfbridgeUnlockAction) ActionTypeName() string {
	return "ConfbridgeUnlock"
}
func (a ConfbridgeUnlockAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeUnlockAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeUnlock(conference string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ConfbridgeUnlockAction{
		Conference: conference,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
