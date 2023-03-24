package actions

type QueuePenaltyAction struct {
	ActionID  string
	Interface string
	Penalty   string
	Queue     string
}

func (QueuePenaltyAction) ActionTypeName() string {
	return "QueuePenalty"
}
func (a QueuePenaltyAction) GetActionID() string {
	return a.ActionID
}
func (a *QueuePenaltyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueuePenalty(iface string, penalty string, opts ...RequestOption) (res *Response, err error) {
	req := &QueuePenaltyAction{
		Interface: iface,
		Penalty:   penalty,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
