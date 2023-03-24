package actions

type PJSIPShowContactsAction struct {
	ActionID string
}

func (PJSIPShowContactsAction) ActionTypeName() string {
	return "PJSIPShowContacts"
}
func (a PJSIPShowContactsAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPShowContactsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) PJSIPShowContacts(opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPShowContactsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
