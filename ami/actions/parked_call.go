package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) ParkedCalls(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ParkedCallsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
