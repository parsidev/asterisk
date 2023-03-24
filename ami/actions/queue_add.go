package actions

type QueueAddAction struct {
	ActionID       string
	Queue          string
	Interface      string
	Penalty        string
	Paused         string
	MemberName     string
	StateInterface string
}

func (QueueAddAction) ActionTypeName() string {
	return "QueueAdd"
}
func (a QueueAddAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueAddAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueAdd(queue string, iface string, opts ...RequestOption) (res *Response, err error) {
	req := &QueueAddAction{
		Queue:     queue,
		Interface: iface,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
