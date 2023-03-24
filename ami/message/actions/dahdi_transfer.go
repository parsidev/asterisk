package actions

import "github.com/parsidev/asterisk/ami/message"

type DAHDITransferAction struct {
	ActionID     string
	DAHDIChannel string
}

func (DAHDITransferAction) ActionTypeName() string {
	return "DAHDITransfer"
}
func (a DAHDITransferAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDITransferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDITransfer(dAHDIChannel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DAHDITransferAction{
		DAHDIChannel: dAHDIChannel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
