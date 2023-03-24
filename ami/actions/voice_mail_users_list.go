package actions

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
func (cli *Client) VoicemailUsersList(opts ...RequestOption) (res *Response, err error) {
	req := &VoicemailUsersListAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
