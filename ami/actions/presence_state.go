package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) PresenceState(provider string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &PresenceStateAction{
		Provider: provider,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
