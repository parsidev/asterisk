package actions

type MixMonitorAction struct {
	ActionID string
	Channel  string
	File     string
	Options  string
	Command  string
}

func (MixMonitorAction) ActionTypeName() string {
	return "MixMonitor"
}
func (a MixMonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *MixMonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MixMonitor(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &MixMonitorAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
