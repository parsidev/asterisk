package actions

import "github.com/parsidev/asterisk/ami/message"

type BridgeTechnologyListAction struct {
	ActionID string
}

func (BridgeTechnologyListAction) ActionTypeName() string {
	return "BridgeTechnologyList"
}
func (a BridgeTechnologyListAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeTechnologyListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) BridgeTechnologyList(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &BridgeTechnologyListAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
