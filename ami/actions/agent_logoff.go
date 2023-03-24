package actions

type AgentLogoffAction struct {
	ActionID string
	Agent    string
	Soft     string
}

func (AgentLogoffAction) ActionTypeName() string {
	return "AgentLogoff"
}
func (a AgentLogoffAction) GetActionID() string {
	return a.ActionID
}
func (a *AgentLogoffAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) AgentLogoff(agent string, opts ...RequestOption) (res *Response, err error) {
	req := &AgentLogoffAction{
		Agent: agent,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
