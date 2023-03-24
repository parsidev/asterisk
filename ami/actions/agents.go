package actions

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
func (cli *Client) Agents(opts ...RequestOption) (res *Response, err error) {
	req := &AgentsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
