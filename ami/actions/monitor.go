package actions

type MonitorAction struct {
	ActionID string
	Channel  string
	File     string
	Format   string
	Mix      string
}

func (MonitorAction) ActionTypeName() string {
	return "Monitor"
}
func (a MonitorAction) GetActionID() string {
	return a.ActionID
}
func (a *MonitorAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Monitor(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &MonitorAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
