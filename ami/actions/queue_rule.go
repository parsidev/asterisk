package actions

import "github.com/parsidev/asterisk/ami/message"

type QueueRuleAction struct {
	ActionID string
	Rule     string
}

func (QueueRuleAction) ActionTypeName() string {
	return "QueueRule"
}
func (a QueueRuleAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueRuleAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueRule(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueRuleAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
