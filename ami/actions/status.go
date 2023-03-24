package actions

type StatusAction struct {
	ActionID     string
	Channel      string
	Variables    string
	AllVariables string
}

func (StatusAction) ActionTypeName() string {
	return "Status"
}
func (a StatusAction) GetActionID() string {
	return a.ActionID
}
func (a *StatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) Status(opts ...RequestOption) (res *Response, err error) {
	req := &StatusAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
