package actions

type PJSIPRegisterAction struct {
	ActionID     string
	Registration string
}

func (PJSIPRegisterAction) ActionTypeName() string {
	return "PJSIPRegister"
}
func (a PJSIPRegisterAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPRegisterAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPRegister(registration string, opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPRegisterAction{
		Registration: registration,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
