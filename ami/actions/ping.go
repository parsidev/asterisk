package actions

import "github.com/parsidev/asterisk/ami/message"

type PingAction struct {
	ActionID string
}

func (PingAction) ActionTypeName() string {
	return "Ping"
}
func (a PingAction) GetActionID() string {
	return a.ActionID
}
func (a *PingAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Ping(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PingAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
