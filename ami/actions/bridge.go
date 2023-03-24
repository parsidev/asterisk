package actions

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
func (cli *Client) Bridge(channel1 string, channel2 string, opts ...RequestOption) (res *Response, err error) {
	req := &BridgeAction{
		Channel1: channel1,
		Channel2: channel2,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
