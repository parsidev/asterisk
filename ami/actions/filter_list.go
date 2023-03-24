package actions

type FilterListAction struct {
	ActionID string
}

func (FilterListAction) ActionTypeName() string {
	return "FilterList"
}
func (a FilterListAction) GetActionID() string {
	return a.ActionID
}
func (a *FilterListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) FilterList(opts ...RequestOption) (res *Response, err error) {
	req := &FilterListAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
