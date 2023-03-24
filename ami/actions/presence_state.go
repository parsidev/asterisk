package actions

type PresenceStateAction struct {
	ActionID string
	Provider string
}

func (PresenceStateAction) ActionTypeName() string {
	return "PresenceState"
}
func (a PresenceStateAction) GetActionID() string {
	return a.ActionID
}
func (a *PresenceStateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PresenceState(provider string, opts ...RequestOption) (res *Response, err error) {
	req := &PresenceStateAction{
		Provider: provider,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
