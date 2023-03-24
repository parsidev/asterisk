package actions

type FilterAction struct {
	ActionID  string
	Operation string
	Filter    string
}

func (FilterAction) ActionTypeName() string {
	return "Filter"
}
func (a FilterAction) GetActionID() string {
	return a.ActionID
}
func (a *FilterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Filter(opts ...RequestOption) (res *Response, err error) {
	req := &FilterAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
