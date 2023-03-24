package actions

import "github.com/parsidev/asterisk/ami/message"

type FAXStatsAction struct {
	ActionID string
}

func (FAXStatsAction) ActionTypeName() string {
	return "FAXStats"
}
func (a FAXStatsAction) GetActionID() string {
	return a.ActionID
}
func (a *FAXStatsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) FAXStats(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &FAXStatsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
