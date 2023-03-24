package actions

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
	opts ...RequestOption) (res *Response, err error) {
	req := &ConfbridgeSetSingleVideoSrcAction{
		Conference: conference,
		Channel:    channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
