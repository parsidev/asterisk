package actions

type SendTextAction struct {
	ActionID    string
	Channel     string
	Message     string
	ContentType string
}

func (SendTextAction) ActionTypeName() string {
	return "SendText"
}
func (a SendTextAction) GetActionID() string {
	return a.ActionID
}
func (a *SendTextAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SendText(channel string, msg string, opts ...RequestOption) (res *Response, err error) {
	req := &SendTextAction{
		Channel: channel,
		Message: msg,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
