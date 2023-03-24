package actions

type DAHDIRestartAction struct {
	ActionID string
}

func (DAHDIRestartAction) ActionTypeName() string {
	return "DAHDIRestart"
}
func (a DAHDIRestartAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIRestartAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIRestart(opts ...RequestOption) (res *Response, err error) {
	req := &DAHDIRestartAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
