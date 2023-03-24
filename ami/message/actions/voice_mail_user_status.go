package actions

import "github.com/parsidev/asterisk/ami/message"

type VoicemailUserStatusAction struct {
	ActionID string
	Context  string
	Mailbox  string
}

func (VoicemailUserStatusAction) ActionTypeName() string {
	return "VoicemailUserStatus"
}
func (a VoicemailUserStatusAction) GetActionID() string {
	return a.ActionID
}
func (a *VoicemailUserStatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) VoicemailUserStatus(context string, mailbox string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &VoicemailUserStatusAction{
		Context: context,
		Mailbox: mailbox,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
