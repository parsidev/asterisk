package actions

type AbsoluteTimeoutAction struct {
	ActionID string
	Channel  string
	Timeout  int
}

func (AbsoluteTimeoutAction) ActionTypeName() string {
	return "AbsoluteTimeout"
}
func (a AbsoluteTimeoutAction) GetActionID() string {
	return a.ActionID
}
func (a *AbsoluteTimeoutAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) AbsoluteTimeout(channel string, timeout int, opts ...RequestOption) (res *Response, err error) {
	req := &AbsoluteTimeoutAction{
		Channel: channel,
		Timeout: timeout,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
