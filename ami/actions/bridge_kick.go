package actions

type BridgeKickAction struct {
	ActionID       string
	BridgeUniqueid string
	Channel        string
}

func (BridgeKickAction) ActionTypeName() string {
	return "BridgeKick"
}
func (a BridgeKickAction) GetActionID() string {
	return a.ActionID
}
func (a *BridgeKickAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) BridgeKick(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &BridgeKickAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
