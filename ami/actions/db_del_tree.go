package actions

type DBDelTreeAction struct {
	ActionID string
	Family   string
	Key      string
}

func (DBDelTreeAction) ActionTypeName() string {
	return "DBDelTree"
}
func (a DBDelTreeAction) GetActionID() string {
	return a.ActionID
}
func (a *DBDelTreeAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DBDelTree(family string, opts ...RequestOption) (res *Response, err error) {
	req := &DBDelTreeAction{
		Family: family,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
