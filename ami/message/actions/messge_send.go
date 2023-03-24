package actions

import "github.com/parsidev/asterisk/ami/message"

type MessageSendAction struct {
	ActionID   string
	To         string
	From       string
	Body       string
	Base64Body string
	Variable   string
}

func (MessageSendAction) ActionTypeName() string {
	return "MessageSend"
}
func (a MessageSendAction) GetActionID() string {
	return a.ActionID
}
func (a *MessageSendAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MessageSend(to string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MessageSendAction{
		To: to,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
