package actions

type IAXpeersAction struct {
	ActionID string
}

func (IAXpeersAction) ActionTypeName() string {
	return "IAXpeers"
}
func (a IAXpeersAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXpeersAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) IAXpeers(opts ...RequestOption) (res *Response, err error) {
	req := &IAXpeersAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
