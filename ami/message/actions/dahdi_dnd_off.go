package actions

import "github.com/parsidev/asterisk/ami/message"

type DAHDIDNDoffAction struct {
	ActionID     string
	DAHDIChannel string
}

func (DAHDIDNDoffAction) ActionTypeName() string {
	return "DAHDIDNDoff"
}
func (a DAHDIDNDoffAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDNDoffAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIDNDoff(dAHDIChannel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DAHDIDNDoffAction{
		DAHDIChannel: dAHDIChannel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
