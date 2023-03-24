package actions

type DAHDIDNDonAction struct {
	ActionID     string
	DAHDIChannel string
}

func (DAHDIDNDonAction) ActionTypeName() string {
	return "DAHDIDNDon"
}
func (a DAHDIDNDonAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDNDonAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIDNDon(dAHDIChannel string, opts ...RequestOption) (res *Response, err error) {
	req := &DAHDIDNDonAction{
		DAHDIChannel: dAHDIChannel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
