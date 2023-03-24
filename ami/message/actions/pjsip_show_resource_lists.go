package actions

import "github.com/parsidev/asterisk/ami/message"

type PJSIPShowResourceListsAction struct {
	ActionID string
}

func (PJSIPShowResourceListsAction) ActionTypeName() string {
	return "PJSIPShowResourceLists"
}
func (a PJSIPShowResourceListsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowResourceListsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPShowResourceLists(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowResourceListsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
