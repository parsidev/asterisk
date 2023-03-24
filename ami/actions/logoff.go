package actions

type LogoffAction struct {
	ActionID string
}

func (LogoffAction) ActionTypeName() string {
	return "Logoff"
}
func (a LogoffAction) GetActionID() string {
	return a.ActionID
}
func (a *LogoffAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Logoff(opts ...RequestOption) (res *Response, err error) {
	req := &LogoffAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
