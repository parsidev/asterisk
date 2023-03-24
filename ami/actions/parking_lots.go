package actions

type ParkinglotsAction struct {
	ActionID string
}

func (ParkinglotsAction) ActionTypeName() string {
	return "Parkinglots"
}
func (a ParkinglotsAction) GetActionID() string {
	return a.ActionID
}
func (a *ParkinglotsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Parkinglots(opts ...RequestOption) (res *Response, err error) {
	req := &ParkinglotsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
