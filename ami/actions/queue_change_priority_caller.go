package actions

type QueueChangePriorityCallerAction struct {
	ActionID string
	Queue    string
	Caller   string
	Priority int
}

func (QueueChangePriorityCallerAction) ActionTypeName() string {
	return "QueueChangePriorityCaller"
}
func (a QueueChangePriorityCallerAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueChangePriorityCallerAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueChangePriorityCaller(queue string, caller string, priority int,
	opts ...RequestOption) (res *Response, err error) {
	req := &QueueChangePriorityCallerAction{
		Queue:    queue,
		Caller:   caller,
		Priority: priority,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
