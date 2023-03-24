package actions

type PJSIPUnregisterAction struct {
	ActionID     string
	Registration string
}

func (PJSIPUnregisterAction) ActionTypeName() string {
	return "PJSIPUnregister"
}
func (a PJSIPUnregisterAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPUnregisterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPUnregister(registration string, opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPUnregisterAction{
		Registration: registration,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
