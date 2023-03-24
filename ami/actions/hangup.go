package actions

type HangupAction struct {
	ActionID string
	Channel  string
	Cause    string
}

func (HangupAction) ActionTypeName() string {
	return "Hangup"
}
func (a HangupAction) GetActionID() string {
	return a.ActionID
}
func (a *HangupAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Hangup(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &HangupAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
