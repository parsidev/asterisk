package actions

import "github.com/parsidev/asterisk/ami/message"

type LoggerRotateAction struct {
	ActionID string
}

func (LoggerRotateAction) ActionTypeName() string {
	return "LoggerRotate"
}
func (a LoggerRotateAction) GetActionID() string {
	return a.ActionID
}
func (a *LoggerRotateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) LoggerRotate(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &LoggerRotateAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
