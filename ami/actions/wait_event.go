package actions

type WaitEventAction struct {
	ActionID string
	Timeout  int
}

func (WaitEventAction) ActionTypeName() string {
	return "WaitEvent"
}
func (a WaitEventAction) GetActionID() string {
	return a.ActionID
}
func (a *WaitEventAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) WaitEvent(timeout int, opts ...RequestOption) (res *Response, err error) {
	req := &WaitEventAction{
		Timeout: timeout,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
