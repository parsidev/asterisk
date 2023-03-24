package actions

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
func (cli *Client) BridgeTechnologyList(opts ...RequestOption) (res *Response, err error) {
	req := &BridgeTechnologyListAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
