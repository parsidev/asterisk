package actions

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
func (cli *Client) BridgeDestroy(bridgeUniqueid string, opts ...RequestOption) (res *Response, err error) {
	req := &BridgeDestroyAction{
		BridgeUniqueid: bridgeUniqueid,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
