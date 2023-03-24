package actions

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

func (cli *Client) DeviceStateList(opts ...RequestOption) (res *Response, err error) {
	req := &DeviceStateListAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
