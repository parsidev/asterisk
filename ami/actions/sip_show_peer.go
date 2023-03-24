package actions

type SIPshowpeerAction struct {
	ActionID string
	Peer     string
}

func (SIPshowpeerAction) ActionTypeName() string {
	return "SIPshowpeer"
}
func (a SIPshowpeerAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPshowpeerAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SIPshowpeer(peer string, opts ...RequestOption) (res *Response, err error) {
	req := &SIPshowpeerAction{
		Peer: peer,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
