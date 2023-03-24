package actions

type DialplanExtensionRemoveAction struct {
	ActionID  string
	Context   string
	Extension string
	Priority  int
}

func (DialplanExtensionRemoveAction) ActionTypeName() string {
	return "DialplanExtensionRemove"
}
func (a DialplanExtensionRemoveAction) GetActionID() string {
	return a.ActionID
}
func (a *DialplanExtensionRemoveAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DialplanExtensionRemove(context string, extension string, opts ...RequestOption) (res *Response, err error) {
	req := &DialplanExtensionRemoveAction{
		Context:   context,
		Extension: extension,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
