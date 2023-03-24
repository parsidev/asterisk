package actions

type MeetmeUnmuteAction struct {
	ActionID string
	Meetme   string
	Usernum  string
}

func (MeetmeUnmuteAction) ActionTypeName() string {
	return "MeetmeUnmute"
}
func (a MeetmeUnmuteAction) GetActionID() string {
	return a.ActionID
}
func (a *MeetmeUnmuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MeetmeUnmute(meetme string, usernum string, opts ...RequestOption) (res *Response, err error) {
	req := &MeetmeUnmuteAction{
		Meetme:  meetme,
		Usernum: usernum,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
