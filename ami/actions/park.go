package actions

import "github.com/parsidev/asterisk/ami/message"

type ParkAction struct {
	ActionID        string
	Channel         string
	TimeoutChannel  string
	AnnounceChannel string
	Timeout         int
	Parkinglot      string
}

func (ParkAction) ActionTypeName() string {
	return "Park"
}
func (a ParkAction) GetActionID() string {
	return a.ActionID
}
func (a *ParkAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Park(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ParkAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
