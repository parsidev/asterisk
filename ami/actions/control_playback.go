package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) ControlPlayback(channel string, control string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ControlPlaybackAction{
		Channel: channel,
		Control: control,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
