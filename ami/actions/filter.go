package actions

import "github.com/parsidev/asterisk/ami/message"

type FilterAction struct {
	ActionID  string
	Operation string
	Filter    string
}

func (FilterAction) ActionTypeName() string {
	return "Filter"
}
func (a FilterAction) GetActionID() string {
	return a.ActionID
}
func (a *FilterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Filter(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &FilterAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
