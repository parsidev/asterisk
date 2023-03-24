package actions

import "github.com/parsidev/asterisk/ami/message"

type VoicemailUsersListAction struct {
	ActionID string
}

func (VoicemailUsersListAction) ActionTypeName() string {
	return "VoicemailUsersList"
}
func (a VoicemailUsersListAction) GetActionID() string {
	return a.ActionID
}
func (a *VoicemailUsersListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) VoicemailUsersList(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &VoicemailUsersListAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
