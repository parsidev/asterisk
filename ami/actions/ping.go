package actions

type PingAction struct {
	ActionID string
}

func (PingAction) ActionTypeName() string {
	return "Ping"
}
func (a PingAction) GetActionID() string {
	return a.ActionID
}
func (a *PingAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Ping(opts ...RequestOption) (res *Response, err error) {
	req := &PingAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
