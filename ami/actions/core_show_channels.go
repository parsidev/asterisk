package actions

type CoreShowChannelsAction struct {
	ActionID string
}

func (CoreShowChannelsAction) ActionTypeName() string {
	return "CoreShowChannels"
}
func (a CoreShowChannelsAction) GetActionID() string {
	return a.ActionID
}
func (a *CoreShowChannelsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}

func (cli *Client) CoreShowChannels(opts ...RequestOption) (res *Response, err error) {
	req := &CoreShowChannelsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
