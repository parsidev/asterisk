package actions

type IAXregistryAction struct {
	ActionID string
}

func (IAXregistryAction) ActionTypeName() string {
	return "IAXregistry"
}
func (a IAXregistryAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXregistryAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) IAXregistry(opts ...RequestOption) (res *Response, err error) {
	req := &IAXregistryAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
