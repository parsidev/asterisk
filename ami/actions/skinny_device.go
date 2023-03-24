package actions

type SKINNYdevicesAction struct {
	ActionID string
}

func (SKINNYdevicesAction) ActionTypeName() string {
	return "SKINNYdevices"
}
func (a SKINNYdevicesAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYdevicesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SKINNYdevices(opts ...RequestOption) (res *Response, err error) {
	req := &SKINNYdevicesAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
