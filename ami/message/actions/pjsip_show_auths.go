package actions

import "github.com/parsidev/asterisk/ami/message"

type PJSIPShowAuthsAction struct {
	ActionID string
}

func (PJSIPShowAuthsAction) ActionTypeName() string {
	return "PJSIPShowAuths"
}
func (a PJSIPShowAuthsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowAuthsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PJSIPShowAuths(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPShowAuthsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
