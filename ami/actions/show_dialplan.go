package actions

import "github.com/parsidev/asterisk/ami/message"

type ShowDialPlanAction struct {
	ActionID  string
	Extension string
	Context   string
}

func (ShowDialPlanAction) ActionTypeName() string {
	return "ShowDialPlan"
}
func (a ShowDialPlanAction) GetActionID() string {
	return a.ActionID
}
func (a *ShowDialPlanAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ShowDialPlan(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ShowDialPlanAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
