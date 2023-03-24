package actions

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

func (cli *Client) BridgeInfo(bridgeUniqueId string, opts ...RequestOption) (res *Response, err error) {
	req := &BridgeInfoAction{
		BridgeUniqueid: bridgeUniqueId,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
