package actions

type LoggerRotateAction struct {
	ActionID string
}

func (LoggerRotateAction) ActionTypeName() string {
	return "LoggerRotate"
}
func (a LoggerRotateAction) GetActionID() string {
	return a.ActionID
}
func (a *LoggerRotateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) LoggerRotate(opts ...RequestOption) (res *Response, err error) {
	req := &LoggerRotateAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
