package actions

type LoginAction struct {
	ActionID string
	UserName string
	Secret   string
}

func (LoginAction) ActionTypeName() string {
	return "Login"
}
func (a LoginAction) GetActionID() string {
	return a.ActionID
}
func (a *LoginAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Login(userName string, opts ...RequestOption) (res *Response, err error) {
	req := &LoginAction{
		UserName: userName,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
