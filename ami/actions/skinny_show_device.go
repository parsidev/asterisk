package actions

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
func (cli *Client) SKINNYshowdevice(device string, opts ...RequestOption) (res *Response, err error) {
	req := &SKINNYshowdeviceAction{
		Device: device,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
