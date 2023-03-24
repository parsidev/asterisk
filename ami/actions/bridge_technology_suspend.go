package actions

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
func (cli *Client) BridgeTechnologySuspend(bridgeTechnology string, opts ...RequestOption) (res *Response, err error) {
	req := &BridgeTechnologySuspendAction{
		BridgeTechnology: bridgeTechnology,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
