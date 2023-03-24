package actions

type MeetmeListRoomsAction struct {
	ActionID string
}

func (MeetmeListRoomsAction) ActionTypeName() string {
	return "MeetmeListRooms"
}
func (a MeetmeListRoomsAction) GetActionID() string {
	return a.ActionID
}
func (a *MeetmeListRoomsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MeetmeListRooms(opts ...RequestOption) (res *Response, err error) {
	req := &MeetmeListRoomsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
