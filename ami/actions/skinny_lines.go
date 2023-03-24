package actions

type SKINNYlinesAction struct {
	ActionID string
}

func (SKINNYlinesAction) ActionTypeName() string {
	return "SKINNYlines"
}
func (a SKINNYlinesAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYlinesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SKINNYlines(opts ...RequestOption) (res *Response, err error) {
	req := &SKINNYlinesAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
