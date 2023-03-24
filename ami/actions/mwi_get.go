package actions

type MWIGetAction struct {
	ActionID string
	Mailbox  string
}

func (MWIGetAction) ActionTypeName() string {
	return "MWIGet"
}
func (a MWIGetAction) GetActionID() string {
	return a.ActionID
}
func (a *MWIGetAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MWIGet(mailbox string, opts ...RequestOption) (res *Response, err error) {
	req := &MWIGetAction{
		Mailbox: mailbox,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
