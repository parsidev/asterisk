package actions

type QueueRemoveAction struct {
	ActionID  string
	Queue     string
	Interface string
}

func (QueueRemoveAction) ActionTypeName() string {
	return "QueueRemove"
}
func (a QueueRemoveAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueRemoveAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueRemove(queue string, iface string, opts ...RequestOption) (res *Response, err error) {
	req := &QueueRemoveAction{
		Queue:     queue,
		Interface: iface,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
