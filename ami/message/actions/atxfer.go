package actions

import "github.com/parsidev/asterisk/ami/message"

type AtxferAction struct {
	ActionID string
	Channel  string
	Exten    string
	Context  string
}

func (AtxferAction) ActionTypeName() string {
	return "Atxfer"
}
func (a AtxferAction) GetActionID() string {
	return a.ActionID
}
func (a *AtxferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Atxfer(channel string, exten string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &AtxferAction{
		Channel: channel,
		Exten:   exten,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
