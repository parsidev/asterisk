package actions

type UserEventAction struct {
	ActionID  string
	UserEvent string
	Header1   string
	HeaderN   string
}

func (UserEventAction) ActionTypeName() string {
	return "UserEvent"
}
func (a UserEventAction) GetActionID() string {
	return a.ActionID
}
func (a *UserEventAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) UserEvent(userEvent string, opts ...RequestOption) (res *Response, err error) {
	req := &UserEventAction{
		UserEvent: userEvent,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
