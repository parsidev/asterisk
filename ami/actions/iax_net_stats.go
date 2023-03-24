package actions

type IAXnetstatsAction struct {
	ActionID string
}

func (IAXnetstatsAction) ActionTypeName() string {
	return "IAXnetstats"
}
func (a IAXnetstatsAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXnetstatsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) IAXnetstats(opts ...RequestOption) (res *Response, err error) {
	req := &IAXnetstatsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
