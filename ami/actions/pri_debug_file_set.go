package actions

type PRIDebugFileSetAction struct {
	ActionID string
	File     string
}

func (PRIDebugFileSetAction) ActionTypeName() string {
	return "PRIDebugFileSet"
}
func (a PRIDebugFileSetAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIDebugFileSetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PRIDebugFileSet(file string, opts ...RequestOption) (res *Response, err error) {
	req := &PRIDebugFileSetAction{
		File: file,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
