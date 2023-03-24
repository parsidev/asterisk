package actions

import "github.com/parsidev/asterisk/ami/message"

type MailboxCountAction struct {
	ActionID string
	Mailbox  string
}

func (MailboxCountAction) ActionTypeName() string {
	return "MailboxCount"
}
func (a MailboxCountAction) GetActionID() string {
	return a.ActionID
}
func (a *MailboxCountAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MailboxCount(mailbox string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MailboxCountAction{
		Mailbox: mailbox,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
