package actions

type ConfbridgeListAction struct {
	ActionID   string
	Conference string
}

func (ConfbridgeListAction) ActionTypeName() string {
	return "ConfbridgeList"
}
func (a ConfbridgeListAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeList(conference string, opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeListAction{
		Conference: conference,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
