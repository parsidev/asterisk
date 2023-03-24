package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) MeetmeUnmute(meetme string, usernum string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &MeetmeUnmuteAction{
		Meetme:  meetme,
		Usernum: usernum,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
