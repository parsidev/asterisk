package actions

type PJSIPNotifyAction struct {
	ActionID string
	Endpoint string
	URI      string
	Channel  string
	Variable string
}

func (PJSIPNotifyAction) ActionTypeName() string {
	return "PJSIPNotify"
}
func (a PJSIPNotifyAction) GetActionID() string {
	return a.ActionID
}
func (a *PJSIPNotifyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PJSIPNotify(variable string, opts ...RequestOption) (res *Response, err error) {
	req := &PJSIPNotifyAction{
		Variable: variable,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
