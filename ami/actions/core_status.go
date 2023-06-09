package actions

type CoreStatusAction struct {
	ActionID string
}

func (CoreStatusAction) ActionTypeName() string {
	return "CoreStatus"
}
func (a CoreStatusAction) GetActionID() string {
	return a.ActionID
}
func (a *CoreStatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) CoreStatus(opts ...RequestOption) (res *Response, err error) {
	req := &CoreStatusAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
