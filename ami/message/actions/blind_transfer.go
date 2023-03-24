package actions

import "github.com/parsidev/asterisk/ami/message"

type BlindTransferAction struct {
	ActionID string
	Channel  string
	Context  string
	Exten    string
}

func (BlindTransferAction) ActionTypeName() string {
	return "BlindTransfer"
}
func (a BlindTransferAction) GetActionID() string {
	return a.ActionID
}
func (a *BlindTransferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) BlindTransfer(channel string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &BlindTransferAction{
		Channel: channel,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
