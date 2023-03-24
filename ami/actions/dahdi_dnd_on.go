package actions

import "github.com/parsidev/asterisk/ami/message"

type DAHDIDNDonAction struct {
	ActionID     string
	DAHDIChannel string
}

func (DAHDIDNDonAction) ActionTypeName() string {
	return "DAHDIDNDon"
}
func (a DAHDIDNDonAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDNDonAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIDNDon(dAHDIChannel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DAHDIDNDonAction{
		DAHDIChannel: dAHDIChannel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
