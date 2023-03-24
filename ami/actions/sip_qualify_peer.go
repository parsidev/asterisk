package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) SIPqualifypeer(peer string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SIPqualifypeerAction{
		Peer: peer,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
