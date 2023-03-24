package actions

import "github.com/parsidev/asterisk/ami/message"

type JabberSendAction struct {
	ActionID string
	Jabber   string
	JID      string
	Message  string
}

func (JabberSendAction) ActionTypeName() string {
	return "JabberSend"
}
func (a JabberSendAction) GetActionID() string {
	return a.ActionID
}
func (a *JabberSendAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) JabberSend(jabber string, jID string, msg string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &JabberSendAction{
		Jabber:  jabber,
		JID:     jID,
		Message: msg,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
