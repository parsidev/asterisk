package actions

import "github.com/parsidev/asterisk/ami/message"

type BridgeTechnologySuspendAction struct {
	ActionID         string
	BridgeTechnology string
}

func (BridgeTechnologySuspendAction) ActionTypeName() string {
	return "BridgeTechnologySuspend"
}
func (a BridgeTechnologySuspendAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeTechnologySuspendAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) BridgeTechnologySuspend(bridgeTechnology string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &BridgeTechnologySuspendAction{
		BridgeTechnology: bridgeTechnology,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
