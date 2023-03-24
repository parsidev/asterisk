package actions

type SIPshowregistryAction struct {
	ActionID string
}

func (SIPshowregistryAction) ActionTypeName() string {
	return "SIPshowregistry"
}
func (a SIPshowregistryAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPshowregistryAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SIPshowregistry(opts ...RequestOption) (res *Response, err error) {
	req := &SIPshowregistryAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
