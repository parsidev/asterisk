package actions

type EventsAction struct {
	ActionID  string
	EventMask string
}

func (EventsAction) ActionTypeName() string {
	return "Events"
}
func (a EventsAction) GetActionID() string {
	return a.ActionID
}
func (a *EventsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Events(eventMask string, opts ...RequestOption) (res *Response, err error) {
	req := &EventsAction{
		EventMask: eventMask,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
