package actions

import "github.com/parsidev/asterisk/ami/message"

type AGIAction struct {
	ActionID  string
	Channel   string
	Command   string
	CommandID string
}

func (AGIAction) ActionTypeName() string {
	return "AGI"
}
func (a AGIAction) GetActionID() string {
	return a.ActionID
}
func (a *AGIAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) AGI(channel string, command string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &AGIAction{
		Channel: channel,
		Command: command,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
