package actions

type ChallengeAction struct {
	ActionID string
	AuthType string
}

func (ChallengeAction) ActionTypeName() string {
	return "Challenge"
}
func (a ChallengeAction) GetActionID() string {
	return a.ActionID
}
func (a *ChallengeAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Challenge(authType string, opts ...RequestOption) (res *Response, err error) {
	req := &ChallengeAction{
		AuthType: authType,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
