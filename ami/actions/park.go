package actions

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
func (cli *Client) Park(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &ParkAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
