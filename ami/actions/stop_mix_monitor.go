package actions

type StopMixMonitorAction struct {
	ActionID     string
	Channel      string
	MixMonitorID string
}

func (StopMixMonitorAction) ActionTypeName() string {
	return "StopMixMonitor"
}
func (a StopMixMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *StopMixMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) StopMixMonitor(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &StopMixMonitorAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
