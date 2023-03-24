package actions

import "github.com/parsidev/asterisk/ami/message"

type IAXpeersAction struct {
	ActionID string
}

func (IAXpeersAction) ActionTypeName() string {
	return "IAXpeers"
}
func (a IAXpeersAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXpeersAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) IAXpeers(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &IAXpeersAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
