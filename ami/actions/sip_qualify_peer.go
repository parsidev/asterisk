package actions

type SIPqualifypeerAction struct {
	ActionID string
	Peer     string
}

func (SIPqualifypeerAction) ActionTypeName() string {
	return "SIPqualifypeer"
}
func (a SIPqualifypeerAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPqualifypeerAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SIPqualifypeer(peer string, opts ...RequestOption) (res *Response, err error) {
	req := &SIPqualifypeerAction{
		Peer: peer,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
