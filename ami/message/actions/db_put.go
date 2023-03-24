package actions

import "github.com/parsidev/asterisk/ami/message"

type DBPutAction struct {
	ActionID string
	Family   string
	Key      string
	Val      string
}

func (DBPutAction) ActionTypeName() string {
	return "DBPut"
}
func (a DBPutAction) GetActionID() string {
	return a.ActionID
}
func (a *DBPutAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DBPut(family string, key string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DBPutAction{
		Family: family,
		Key:    key,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
