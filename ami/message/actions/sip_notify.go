package actions

import "github.com/parsidev/asterisk/ami/message"

type SIPnotifyAction struct {
	ActionID string
	Channel  string
	Variable string
	CallID   string
}

func (SIPnotifyAction) ActionTypeName() string {
	return "SIPnotify"
}
func (a SIPnotifyAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPnotifyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SIPnotify(channel string, variable string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SIPnotifyAction{
		Channel:  channel,
		Variable: variable,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
