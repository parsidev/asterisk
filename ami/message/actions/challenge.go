package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) Challenge(authType string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ChallengeAction{
		AuthType: authType,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
