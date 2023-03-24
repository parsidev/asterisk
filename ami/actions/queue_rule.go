package actions

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
func (cli *Client) QueueRule(opts ...RequestOption) (res *Response, err error) {
	req := &QueueRuleAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
