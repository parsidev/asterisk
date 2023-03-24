package actions

import "github.com/parsidev/asterisk/ami/message"

type PJSIPShowContactsAction struct {
	ActionID string
}

func (PJSIPShowContactsAction) ActionTypeName() string {
	return "PJSIPShowContacts"
}
func (a PJSIPShowContactsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowContactsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PJSIPShowContacts(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowContactsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
