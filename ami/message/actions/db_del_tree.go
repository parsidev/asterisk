package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) DBDelTree(family string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DBDelTreeAction{
		Family: family,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
