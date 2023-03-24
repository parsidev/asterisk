package actions

type SetvarAction struct {
	ActionID string
	Channel  string
	Variable string
	Value    string
}

func (SetvarAction) ActionTypeName() string {
	return "Setvar"
}
func (a SetvarAction) GetActionID() string {
	return a.ActionID
}
func (a *SetvarAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Setvar(variable string, value string, opts ...RequestOption) (res *Response, err error) {
	req := &SetvarAction{
		Variable: variable,
		Value:    value,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
