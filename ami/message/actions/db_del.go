package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) DBDel(family string, key string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DBDelAction{
		Family: family,
		Key:    key,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
