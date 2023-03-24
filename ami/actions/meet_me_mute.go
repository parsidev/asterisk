package actions

type MeetmeMuteAction struct {
	ActionID string
	Meetme   string
	Usernum  string
}

func (MeetmeMuteAction) ActionTypeName() string {
	return "MeetmeMute"
}
func (a MeetmeMuteAction) GetActionID() string {
	return a.ActionID
}
func (a *MeetmeMuteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MeetmeMute(meetme string, usernum string, opts ...RequestOption) (res *Response, err error) {
	req := &MeetmeMuteAction{
		Meetme:  meetme,
		Usernum: usernum,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
