package actions

type MWIDeleteAction struct {
	ActionID string
	Mailbox  string
}

func (MWIDeleteAction) ActionTypeName() string {
	return "MWIDelete"
}
func (a MWIDeleteAction) GetActionID() string {
	return a.ActionID
}
func (a *MWIDeleteAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MWIDelete(mailbox string, opts ...RequestOption) (res *Response, err error) {
	req := &MWIDeleteAction{
		Mailbox: mailbox,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
