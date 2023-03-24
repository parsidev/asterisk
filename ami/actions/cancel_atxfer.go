package actions

type CancelAtxferAction struct {
	ActionID string
	Channel  string
}

func (CancelAtxferAction) ActionTypeName() string {
	return "CancelAtxfer"
}
func (a CancelAtxferAction) GetActionID() string {
	return a.ActionID
}
func (a *CancelAtxferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) CancelAtxfer(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &CancelAtxferAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
