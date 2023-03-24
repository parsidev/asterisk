package actions

import "github.com/parsidev/asterisk/ami/message"

type MailboxStatusAction struct {
	ActionID string
	Mailbox  string
}

func (MailboxStatusAction) ActionTypeName() string {
	return "MailboxStatus"
}
func (a MailboxStatusAction) GetActionID() string {
	return a.ActionID
}
func (a *MailboxStatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MailboxStatus(mailbox string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MailboxStatusAction{
		Mailbox: mailbox,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
