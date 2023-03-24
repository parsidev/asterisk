package actions

type ConfbridgeLockAction struct {
	ActionID   string
	Conference string
}

func (ConfbridgeLockAction) ActionTypeName() string {
	return "ConfbridgeLock"
}
func (a ConfbridgeLockAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeLockAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeLock(conference string, opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeLockAction{
		Conference: conference,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
