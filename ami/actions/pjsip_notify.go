package actions

import "github.com/parsidev/asterisk/ami/message"

type PJSIPNotifyAction struct {
	ActionID string
	Endpoint string
	URI      string
	Channel  string
	Variable string
}

func (PJSIPNotifyAction) ActionTypeName() string {
	return "PJSIPNotify"
}
func (a PJSIPNotifyAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPNotifyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPNotify(variable string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPNotifyAction{
		Variable: variable,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
