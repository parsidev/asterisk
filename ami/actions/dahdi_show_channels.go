package actions

type DAHDIShowChannelsAction struct {
	ActionID     string
	DAHDIChannel string
}

func (DAHDIShowChannelsAction) ActionTypeName() string {
	return "DAHDIShowChannels"
}
func (a DAHDIShowChannelsAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIShowChannelsAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIShowChannels(opts ...RequestOption) (res *Response, err error) {
	req := &DAHDIShowChannelsAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
