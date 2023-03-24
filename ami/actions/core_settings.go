package actions

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
func (cli *Client) CoreSettings(opts ...RequestOption) (res *Response, err error) {
	req := &CoreSettingsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
