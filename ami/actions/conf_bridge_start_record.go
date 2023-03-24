package actions

type ConfbridgeStartRecordAction struct {
	ActionID   string
	Conference string
	RecordFile string
}

func (ConfbridgeStartRecordAction) ActionTypeName() string {
	return "ConfbridgeStartRecord"
}
func (a ConfbridgeStartRecordAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeStartRecordAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeStartRecord(conference string, opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeStartRecordAction{
		Conference: conference,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
