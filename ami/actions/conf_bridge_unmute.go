package actions

type ConfbridgeUnmuteAction struct {
	ActionID   string
	Conference string
	Channel    string
}

func (ConfbridgeUnmuteAction) ActionTypeName() string {
	return "ConfbridgeUnmute"
}
func (a ConfbridgeUnmuteAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeUnmuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeUnmute(conference string, channel string, opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeUnmuteAction{
		Conference: conference,
		Channel:    channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
