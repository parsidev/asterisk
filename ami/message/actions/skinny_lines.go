package actions

import "github.com/parsidev/asterisk/ami/message"

type SKINNYlinesAction struct {
	ActionID string
}

func (SKINNYlinesAction) ActionTypeName() string {
	return "SKINNYlines"
}
func (a SKINNYlinesAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYlinesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SKINNYlines(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SKINNYlinesAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
