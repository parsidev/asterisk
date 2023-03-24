package actions

import "github.com/parsidev/asterisk/ami/message"

type PlayDTMFAction struct {
	ActionID string
	Channel  string
	Digit    string
	Duration string
	Receive  string
}

func (PlayDTMFAction) ActionTypeName() string {
	return "PlayDTMF"
}
func (a PlayDTMFAction) GetActionID() string {
	return a.ActionID
}
func (a *PlayDTMFAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PlayDTMF(channel string, digit string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PlayDTMFAction{
		Channel: channel,
		Digit:   digit,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
