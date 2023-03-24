package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) SIPshowpeer(peer string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SIPshowpeerAction{
		Peer: peer,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
