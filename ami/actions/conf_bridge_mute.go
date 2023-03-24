package actions

type ConfbridgeMuteAction struct {
	ActionID   string
	Conference string
	Channel    string
}

func (ConfbridgeMuteAction) ActionTypeName() string {
	return "ConfbridgeMute"
}
func (a ConfbridgeMuteAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeMuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeMute(conference string, channel string, opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeMuteAction{
		Conference: conference,
		Channel:    channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
