package actions

type MeetmeListAction struct {
	ActionID   string
	Conference string
}

func (MeetmeListAction) ActionTypeName() string {
	return "MeetmeList"
}
func (a MeetmeListAction) GetActionID() string {
	return a.ActionID
}
func (a *MeetmeListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MeetmeList(opts ...RequestOption) (res *Response, err error) {
	req := &MeetmeListAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
