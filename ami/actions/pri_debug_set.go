package actions

type PRIDebugSetAction struct {
	ActionID string
	Span     string
	Level    string
}

func (PRIDebugSetAction) ActionTypeName() string {
	return "PRIDebugSet"
}
func (a PRIDebugSetAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIDebugSetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PRIDebugSet(span string, level string, opts ...RequestOption) (res *Response, err error) {
	req := &PRIDebugSetAction{
		Span:  span,
		Level: level,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
