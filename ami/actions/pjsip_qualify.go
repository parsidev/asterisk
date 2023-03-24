package actions

import "github.com/parsidev/asterisk/ami/message"

type PJSIPQualifyAction struct {
	ActionID string
	Endpoint string
}

func (PJSIPQualifyAction) ActionTypeName() string {
	return "PJSIPQualify"
}
func (a PJSIPQualifyAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPQualifyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPQualify(endpoint string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PJSIPQualifyAction{
		Endpoint: endpoint,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
