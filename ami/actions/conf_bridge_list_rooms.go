package actions

type ConfbridgeListRoomsAction struct {
	ActionID string
}

func (ConfbridgeListRoomsAction) ActionTypeName() string {
	return "ConfbridgeListRooms"
}
func (a ConfbridgeListRoomsAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeListRoomsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeListRooms(opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeListRoomsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
