package actions

type DialplanExtensionAddAction struct {
	ActionID        string
	Context         string
	Extension       string
	Priority        int
	Application     string
	ApplicationData string
	Replace         string
}

func (DialplanExtensionAddAction) ActionTypeName() string {
	return "DialplanExtensionAdd"
}
func (a DialplanExtensionAddAction) GetActionID() string {
	return a.ActionID
}
func (a *DialplanExtensionAddAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DialplanExtensionAdd(context string, extension string, priority int, application string,
	opts ...RequestOption) (res *Response, err error) {
	req := &DialplanExtensionAddAction{
		Context:     context,
		Extension:   extension,
		Priority:    priority,
		Application: application,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
