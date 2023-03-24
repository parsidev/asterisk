package actions

import "github.com/parsidev/asterisk/ami/message"

type BridgeListAction struct {
	ActionID   string
	BridgeType string
}

func (BridgeListAction) ActionTypeName() string {
	return "BridgeList"
}
func (a BridgeListAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) BridgeList(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &BridgeListAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
