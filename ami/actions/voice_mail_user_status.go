package actions

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
func (cli *Client) VoicemailUserStatus(context string, mailbox string, opts ...RequestOption) (res *Response, err error) {
	req := &VoicemailUserStatusAction{
		Context: context,
		Mailbox: mailbox,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
