package actions

type SIPpeersAction struct {
	ActionID string
}

func (SIPpeersAction) ActionTypeName() string {
	return "SIPpeers"
}
func (a SIPpeersAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPpeersAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SIPpeers(opts ...RequestOption) (res *Response, err error) {
	req := &SIPpeersAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
