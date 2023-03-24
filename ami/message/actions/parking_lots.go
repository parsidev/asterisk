package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) Parkinglots(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ParkinglotsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
