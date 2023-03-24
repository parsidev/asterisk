package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) AbsoluteTimeout(channel string, timeout int, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &AbsoluteTimeoutAction{
		Channel: channel,
		Timeout: timeout,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
