package actions

type BridgeTechnologyUnsuspendAction struct {
	ActionID         string
	BridgeTechnology string
}

func (BridgeTechnologyUnsuspendAction) ActionTypeName() string {
	return "BridgeTechnologyUnsuspend"
}
func (a BridgeTechnologyUnsuspendAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeTechnologyUnsuspendAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) BridgeTechnologyUnsuspend(bridgeTechnology string, opts ...RequestOption) (res *Response, err error) {
	req := &BridgeTechnologyUnsuspendAction{
		BridgeTechnology: bridgeTechnology,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
