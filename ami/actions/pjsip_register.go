package actions

import "github.com/parsidev/asterisk/ami/message"

type PJSIPRegisterAction struct {
	ActionID     string
	Registration string
}

func (PJSIPRegisterAction) ActionTypeName() string {
	return "PJSIPRegister"
}
func (a PJSIPRegisterAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPRegisterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPRegister(registration string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPRegisterAction{
		Registration: registration,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
