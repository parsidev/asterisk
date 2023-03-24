package actions

type ConfbridgeKickAction struct {
	ActionID   string
	Conference string
	Channel    string
}

func (ConfbridgeKickAction) ActionTypeName() string {
	return "ConfbridgeKick"
}
func (a ConfbridgeKickAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeKickAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeKick(conference string, channel string, opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeKickAction{
		Conference: conference,
		Channel:    channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
