package actions

import "github.com/parsidev/asterisk/ami/message"

type LogoffAction struct {
	ActionID string
}

func (LogoffAction) ActionTypeName() string {
	return "Logoff"
}
func (a LogoffAction) GetActionID() string {
	return a.ActionID
}
func (a *LogoffAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Logoff(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &LogoffAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
