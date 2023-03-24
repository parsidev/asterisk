package actions

import "github.com/parsidev/asterisk/ami/message"

type BridgeInfoAction struct {
	ActionID       string
	BridgeUniqueid string
}

func (BridgeInfoAction) ActionTypeName() string {
	return "BridgeInfo"
}
func (a BridgeInfoAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeInfoAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) BridgeInfo(bridgeUniqueId string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &BridgeInfoAction{
		BridgeUniqueid: bridgeUniqueId,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
