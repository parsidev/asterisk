package actions

import "github.com/parsidev/asterisk/ami/message"

type DeviceStateListAction struct {
	ActionID string
}

func (DeviceStateListAction) ActionTypeName() string {
	return "DeviceStateList"
}
func (a DeviceStateListAction) GetActionID() string {
	return a.ActionID
}
func (a *DeviceStateListAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) DeviceStateList(opts ...message.RequestOption) (res *message.Response, err error) {
	req := &DeviceStateListAction{}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
