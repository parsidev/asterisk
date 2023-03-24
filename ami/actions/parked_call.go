package actions

type ParkedCallsAction struct {
	ActionID   string
	ParkingLot string
}

func (ParkedCallsAction) ActionTypeName() string {
	return "ParkedCalls"
}
func (a ParkedCallsAction) GetActionID() string {
	return a.ActionID
}
func (a *ParkedCallsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ParkedCalls(opts ...RequestOption) (res *Response, err error) {
	req := &ParkedCallsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
