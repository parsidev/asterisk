package actions

import "github.com/parsidev/asterisk/ami/message"

type CoreSettingsAction struct {
	ActionID string
}

func (CoreSettingsAction) ActionTypeName() string {
	return "CoreSettings"
}
func (a CoreSettingsAction) GetActionID() string {
	return a.ActionID
}
func (a *CoreSettingsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) CoreSettings(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &CoreSettingsAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
