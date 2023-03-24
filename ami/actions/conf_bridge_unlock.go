package actions

type ConfbridgeUnlockAction struct {
	ActionID   string
	Conference string
}

func (ConfbridgeUnlockAction) ActionTypeName() string {
	return "ConfbridgeUnlock"
}
func (a ConfbridgeUnlockAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeUnlockAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeUnlock(conference string, opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeUnlockAction{
		Conference: conference,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
