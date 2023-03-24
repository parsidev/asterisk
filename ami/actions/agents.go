package actions

import "github.com/parsidev/asterisk/ami/message"

type AgentsAction struct {
	ActionID string
}

func (AgentsAction) ActionTypeName() string {
	return "Agents"
}
func (a AgentsAction) GetActionID() string {
	return a.ActionID
}
func (a *AgentsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Agents(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &AgentsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
