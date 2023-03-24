package actions

type DBDelAction struct {
	ActionID string
	Family   string
	Key      string
}

func (DBDelAction) ActionTypeName() string {
	return "DBDel"
}
func (a DBDelAction) GetActionID() string {
	return a.ActionID
}
func (a *DBDelAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DBDel(family string, key string, opts ...RequestOption) (res *Response, err error) {
	req := &DBDelAction{
		Family: family,
		Key:    key,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
