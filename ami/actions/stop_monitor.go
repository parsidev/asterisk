package actions

type StopMonitorAction struct {
	ActionID string
	Channel  string
}

func (StopMonitorAction) ActionTypeName() string {
	return "StopMonitor"
}
func (a StopMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *StopMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) StopMonitor(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &StopMonitorAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
