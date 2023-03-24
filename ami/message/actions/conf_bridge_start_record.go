package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) ConfbridgeStartRecord(conference string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ConfbridgeStartRecordAction{
		Conference: conference,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
