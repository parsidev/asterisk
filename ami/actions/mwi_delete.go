package actions

import "github.com/parsidev/asterisk/ami/message"

type MWIDeleteAction struct {
	ActionID string
	Mailbox  string
}

func (MWIDeleteAction) ActionTypeName() string {
	return "MWIDelete"
}
func (a MWIDeleteAction) GetActionID() string {
	return a.ActionID
}
func (a *MWIDeleteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MWIDelete(mailbox string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MWIDeleteAction{
		Mailbox: mailbox,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
