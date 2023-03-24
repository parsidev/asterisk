package actions

import "github.com/parsidev/asterisk/ami/message"

type LocalOptimizeAwayAction struct {
	ActionID string
	Channel  string
}

func (LocalOptimizeAwayAction) ActionTypeName() string {
	return "LocalOptimizeAway"
}
func (a LocalOptimizeAwayAction) GetActionID() string {
	return a.ActionID
}
func (a *LocalOptimizeAwayAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) LocalOptimizeAway(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &LocalOptimizeAwayAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
