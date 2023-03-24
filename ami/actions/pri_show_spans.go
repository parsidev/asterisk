package actions

import "github.com/parsidev/asterisk/ami/message"

type PRIShowSpansAction struct {
	ActionID string
	Span     string
}

func (PRIShowSpansAction) ActionTypeName() string {
	return "PRIShowSpans"
}
func (a PRIShowSpansAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIShowSpansAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PRIShowSpans(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PRIShowSpansAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
