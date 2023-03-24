package actions

import "github.com/parsidev/asterisk/ami/message"

type FilterListAction struct {
	ActionID string
}

func (FilterListAction) ActionTypeName() string {
	return "FilterList"
}
func (a FilterListAction) GetActionID() string {
	return a.ActionID
}
func (a *FilterListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) FilterList(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &FilterListAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
