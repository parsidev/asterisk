package actions

import "github.com/parsidev/asterisk/ami/message"

type BridgeAction struct {
	ActionID string
	Channel1 string
	Channel2 string
	Tone     string
}

func (BridgeAction) ActionTypeName() string {
	return "Bridge"
}
func (a BridgeAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Bridge(channel1 string, channel2 string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &BridgeAction{
		Channel1: channel1,
		Channel2: channel2,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
