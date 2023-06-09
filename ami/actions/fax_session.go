package actions

type FAXSessionAction struct {
	ActionID      string
	SessionNumber string
}

func (FAXSessionAction) ActionTypeName() string {
	return "FAXSession"
}
func (a FAXSessionAction) GetActionID() string {
	return a.ActionID
}
func (a *FAXSessionAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) FAXSession(sessionNumber string, opts ...RequestOption) (res *Response, err error) {
	req := &FAXSessionAction{
		SessionNumber: sessionNumber,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
