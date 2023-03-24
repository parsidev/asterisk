package actions

import "github.com/parsidev/asterisk/ami/message"

type FAXSessionsAction struct {
	ActionID string
}

func (FAXSessionsAction) ActionTypeName() string {
	return "FAXSessions"
}
func (a FAXSessionsAction) GetActionID() string {
	return a.ActionID
}
func (a *FAXSessionsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) FAXSessions(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &FAXSessionsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
