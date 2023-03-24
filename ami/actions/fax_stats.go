package actions

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
func (cli *Client) FAXStats(opts ...RequestOption) (res *Response, err error) {
	req := &FAXStatsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
