package actions

import "github.com/parsidev/asterisk/ami/message"

type BridgeDestroyAction struct {
	ActionID       string
	BridgeUniqueid string
}

func (BridgeDestroyAction) ActionTypeName() string {
	return "BridgeDestroy"
}
func (a BridgeDestroyAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeDestroyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) BridgeDestroy(bridgeUniqueid string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &BridgeDestroyAction{
		BridgeUniqueid: bridgeUniqueid,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
