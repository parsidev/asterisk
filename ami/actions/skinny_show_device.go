package actions

import "github.com/parsidev/asterisk/ami/message"

type SKINNYshowdeviceAction struct {
	ActionID string
	Device   string
}

func (SKINNYshowdeviceAction) ActionTypeName() string {
	return "SKINNYshowdevice"
}
func (a SKINNYshowdeviceAction) GetActionID() string {
	return a.ActionID
}
func (a *SKINNYshowdeviceAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SKINNYshowdevice(device string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SKINNYshowdeviceAction{
		Device: device,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
