package actions

type ConfbridgeStopRecordAction struct {
	ActionID   string
	Conference string
}

func (ConfbridgeStopRecordAction) ActionTypeName() string {
	return "ConfbridgeStopRecord"
}
func (a ConfbridgeStopRecordAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeStopRecordAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeStopRecord(conference string, opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeStopRecordAction{
		Conference: conference,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
