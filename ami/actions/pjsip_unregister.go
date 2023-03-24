package actions

import "github.com/parsidev/asterisk/ami/message"

type PJSIPUnregisterAction struct {
	ActionID     string
	Registration string
}

func (PJSIPUnregisterAction) ActionTypeName() string {
	return "PJSIPUnregister"
}
func (a PJSIPUnregisterAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPUnregisterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPUnregister(registration string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPUnregisterAction{
		Registration: registration,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
