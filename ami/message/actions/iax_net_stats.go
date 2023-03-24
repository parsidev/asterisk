package actions

import "github.com/parsidev/asterisk/ami/message"

type IAXnetstatsAction struct {
	ActionID string
}

func (IAXnetstatsAction) ActionTypeName() string {
	return "IAXnetstats"
}
func (a IAXnetstatsAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXnetstatsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) IAXnetstats(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &IAXnetstatsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
