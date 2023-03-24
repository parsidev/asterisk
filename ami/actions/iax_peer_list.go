package actions

type IAXpeerlistAction struct {
	ActionID string
}

func (IAXpeerlistAction) ActionTypeName() string {
	return "IAXpeerlist"
}
func (a IAXpeerlistAction) GetActionID() string {
	return a.ActionID
}
func (a *IAXpeerlistAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) IAXpeerlist(opts ...RequestOption) (res *Response, err error) {
	req := &IAXpeerlistAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
