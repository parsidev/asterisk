package actions

import "github.com/parsidev/asterisk/ami/message"

type DBGetAction struct {
	ActionID string
	Family   string
	Key      string
}

func (DBGetAction) ActionTypeName() string {
	return "DBGet"
}
func (a DBGetAction) GetActionID() string {
	return a.ActionID
}
func (a *DBGetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DBGet(family string, key string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DBGetAction{
		Family: family,
		Key:    key,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
