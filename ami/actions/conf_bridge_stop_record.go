package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) ConfbridgeStopRecord(conference string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ConfbridgeStopRecordAction{
		Conference: conference,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
