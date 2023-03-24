package actions

type ExtensionStateAction struct {
	ActionID string
	Exten    string
	Context  string
}

func (ExtensionStateAction) ActionTypeName() string {
	return "ExtensionState"
}
func (a ExtensionStateAction) GetActionID() string {
	return a.ActionID
}
func (a *ExtensionStateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ExtensionState(exten string, context string, opts ...RequestOption) (res *Response, err error) {
	req := &ExtensionStateAction{
		Exten:   exten,
		Context: context,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
