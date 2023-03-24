package actions

type QueuePauseAction struct {
	ActionID  string
	Interface string
	Paused    string
	Queue     string
	Reason    string
}

func (QueuePauseAction) ActionTypeName() string {
	return "QueuePause"
}
func (a QueuePauseAction) GetActionID() string {
	return a.ActionID
}
func (a *QueuePauseAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueuePause(iface string, paused string, opts ...RequestOption) (res *Response, err error) {
	req := &QueuePauseAction{
		Interface: iface,
		Paused:    paused,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
