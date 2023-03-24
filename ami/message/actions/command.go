package actions

import "github.com/parsidev/asterisk/ami/message"

type CommandAction struct {
	ActionID string
	Command  string
}

func (CommandAction) ActionTypeName() string {
	return "Command"
}
func (a CommandAction) GetActionID() string {
	return a.ActionID
}
func (a *CommandAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Command(command string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &CommandAction{
		Command: command,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
