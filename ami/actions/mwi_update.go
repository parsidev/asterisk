package actions

type MWIUpdateAction struct {
	ActionID    string
	Mailbox     string
	OldMessages string
	NewMessages string
}

func (MWIUpdateAction) ActionTypeName() string {
	return "MWIUpdate"
}
func (a MWIUpdateAction) GetActionID() string {
	return a.ActionID
}
func (a *MWIUpdateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MWIUpdate(mailbox string, opts ...RequestOption) (res *Response, err error) {
	req := &MWIUpdateAction{
		Mailbox: mailbox,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
