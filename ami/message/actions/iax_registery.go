package actions

import "github.com/parsidev/asterisk/ami/message"

type IAXregistryAction struct {
	ActionID string
}

func (IAXregistryAction) ActionTypeName() string {
	return "IAXregistry"
}
func (a IAXregistryAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXregistryAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) IAXregistry(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &IAXregistryAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
