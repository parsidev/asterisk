package actions

type PRIDebugFileUnsetAction struct {
	ActionID string
}

func (PRIDebugFileUnsetAction) ActionTypeName() string {
	return "PRIDebugFileUnset"
}
func (a PRIDebugFileUnsetAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIDebugFileUnsetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PRIDebugFileUnset(opts ...RequestOption) (res *Response, err error) {
	req := &PRIDebugFileUnsetAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
