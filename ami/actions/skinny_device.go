package actions

import "github.com/parsidev/asterisk/ami/message"

type SKINNYdevicesAction struct {
	ActionID string
}

func (SKINNYdevicesAction) ActionTypeName() string {
	return "SKINNYdevices"
}
func (a SKINNYdevicesAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYdevicesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SKINNYdevices(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SKINNYdevicesAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
