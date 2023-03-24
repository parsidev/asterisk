package actions

type ControlPlaybackAction struct {
	ActionID string
	Channel  string
	Control  string
}

func (ControlPlaybackAction) ActionTypeName() string {
	return "ControlPlayback"
}
func (a ControlPlaybackAction) GetActionID() string {
	return a.ActionID
}
func (a *ControlPlaybackAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ControlPlayback(channel string, control string, opts ...RequestOption) (res *Response, err error) {
	req := &ControlPlaybackAction{
		Channel: channel,
		Control: control,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
