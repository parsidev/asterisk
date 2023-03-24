package actions

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
func (cli *Client) FAXSessions(opts ...RequestOption) (res *Response, err error) {
	req := &FAXSessionsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
