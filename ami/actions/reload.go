package actions

type ReloadAction struct {
	ActionID string
	Module   string
}

func (ReloadAction) ActionTypeName() string {
	return "Reload"
}
func (a ReloadAction) GetActionID() string {
	return a.ActionID
}
func (a *ReloadAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Reload(opts ...RequestOption) (res *Response, err error) {
	req := &ReloadAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
