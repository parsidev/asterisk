package actions

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
func (cli *Client) BridgeList(opts ...RequestOption) (res *Response, err error) {
	req := &BridgeListAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
