package actions

type PresenceStateListAction struct {
	ActionID string
}

func (PresenceStateListAction) ActionTypeName() string {
	return "PresenceStateList"
}
func (a PresenceStateListAction) GetActionID() string {
	return a.ActionID
}
func (a *PresenceStateListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PresenceStateList(opts ...RequestOption) (res *Response, err error) {
	req := &PresenceStateListAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
