package actions

type VoicemailRefreshAction struct {
	ActionID string
	Context  string
	Mailbox  string
}

func (VoicemailRefreshAction) ActionTypeName() string {
	return "VoicemailRefresh"
}
func (a VoicemailRefreshAction) GetActionID() string {
	return a.ActionID
}
func (a *VoicemailRefreshAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) VoicemailRefresh(opts ...RequestOption) (res *Response, err error) {
	req := &VoicemailRefreshAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
