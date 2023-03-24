package actions

type GetvarAction struct {
	ActionID string
	Channel  string
	Variable string
}

func (GetvarAction) ActionTypeName() string {
	return "Getvar"
}
func (a GetvarAction) GetActionID() string {
	return a.ActionID
}
func (a *GetvarAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Getvar(variable string, opts ...RequestOption) (res *Response, err error) {
	req := &GetvarAction{
		Variable: variable,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
