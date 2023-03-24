package actions

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
func (cli *Client) MailboxStatus(mailbox string, opts ...RequestOption) (res *Response, err error) {
	req := &MailboxStatusAction{
		Mailbox: mailbox,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
