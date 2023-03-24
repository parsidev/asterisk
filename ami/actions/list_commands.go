package actions

import "github.com/parsidev/asterisk/ami/message"

type ListCommandsAction struct {
	ActionID string
}

func (ListCommandsAction) ActionTypeName() string {
	return "ListCommands"
}
func (a ListCommandsAction) GetActionID() string {
	return a.ActionID
}
func (a *ListCommandsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ListCommands(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ListCommandsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
