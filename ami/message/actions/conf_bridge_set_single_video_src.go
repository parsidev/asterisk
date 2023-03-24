package actions

import "github.com/parsidev/asterisk/ami/message"

type ConfbridgeSetSingleVideoSrcAction struct {
	ActionID   string
	Conference string
	Channel    string
}

func (ConfbridgeSetSingleVideoSrcAction) ActionTypeName() string {
	return "ConfbridgeSetSingleVideoSrc"
}
func (a ConfbridgeSetSingleVideoSrcAction) GetActionID() string {
	return a.ActionID
}
func (a *ConfbridgeSetSingleVideoSrcAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ConfbridgeSetSingleVideoSrc(conference string, channel string,
	opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ConfbridgeSetSingleVideoSrcAction{
		Conference: conference,
		Channel:    channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
