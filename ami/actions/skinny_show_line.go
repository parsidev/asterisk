package actions

import "github.com/parsidev/asterisk/ami/message"

type SKINNYshowlineAction struct {
	ActionID string
	Line     string
}

func (SKINNYshowlineAction) ActionTypeName() string {
	return "SKINNYshowline"
}
func (a SKINNYshowlineAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYshowlineAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SKINNYshowline(line string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SKINNYshowlineAction{
		Line: line,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
