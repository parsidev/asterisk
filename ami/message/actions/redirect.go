package actions

import "github.com/parsidev/asterisk/ami/message"

type RedirectAction struct {
	ActionID      string
	Channel       string
	ExtraChannel  string
	Exten         string
	ExtraExten    string
	Context       string
	ExtraContext  string
	Priority      int
	ExtraPriority string
}

func (RedirectAction) ActionTypeName() string {
	return "Redirect"
}
func (a RedirectAction) GetActionID() string {
	return a.ActionID
}
func (a *RedirectAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Redirect(channel string, exten string, context string, priority int,
	opts ...message.RequestOption) (res *message.Response, err error) {
	req := &RedirectAction{
		Channel:  channel,
		Exten:    exten,
		Context:  context,
		Priority: priority,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
