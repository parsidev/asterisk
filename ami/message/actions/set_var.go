package actions

import "github.com/parsidev/asterisk/ami/message"

type SetvarAction struct {
	ActionID string
	Channel  string
	Variable string
	Value    string
}

func (SetvarAction) ActionTypeName() string {
	return "Setvar"
}
func (a SetvarAction) GetActionID() string {
	return a.ActionID
}
func (a *SetvarAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Setvar(variable string, value string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SetvarAction{
		Variable: variable,
		Value:    value,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
