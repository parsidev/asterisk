package actions

import "github.com/parsidev/asterisk/ami/message"

type MWIGetAction struct {
	ActionID string
	Mailbox  string
}

func (MWIGetAction) ActionTypeName() string {
	return "MWIGet"
}
func (a MWIGetAction) GetActionID() string {
	return a.ActionID
}
func (a *MWIGetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MWIGet(mailbox string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MWIGetAction{
		Mailbox: mailbox,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
