package actions

type AGIAction struct {
	ActionID  string
	Channel   string
	Command   string
	CommandID string
}

func (AGIAction) ActionTypeName() string {
	return "AGI"
}
func (a AGIAction) GetActionID() string {
	return a.ActionID
}
func (a *AGIAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) AGI(channel string, command string, opts ...RequestOption) (res *Response, err error) {
	req := &AGIAction{
		Channel: channel,
		Command: command,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
