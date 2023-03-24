package actions

type MixMonitorMuteAction struct {
	ActionID  string
	Channel   string
	Direction string
	State     string
}

func (MixMonitorMuteAction) ActionTypeName() string {
	return "MixMonitorMute"
}
func (a MixMonitorMuteAction) GetActionID() string {
	return a.ActionID
}
func (a *MixMonitorMuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MixMonitorMute(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &MixMonitorMuteAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
