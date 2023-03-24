package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) WaitEvent(timeout int, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &WaitEventAction{
		Timeout: timeout,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
