package actions

type SIPpeerstatusAction struct {
	ActionID string
	Peer     string
}

func (SIPpeerstatusAction) ActionTypeName() string {
	return "SIPpeerstatus"
}
func (a SIPpeerstatusAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPpeerstatusAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SIPpeerstatus(opts ...RequestOption) (res *Response, err error) {
	req := &SIPpeerstatusAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
