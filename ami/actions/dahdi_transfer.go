package actions

type DAHDITransferAction struct {
	ActionID     string
	DAHDIChannel string
}

func (DAHDITransferAction) ActionTypeName() string {
	return "DAHDITransfer"
}
func (a DAHDITransferAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDITransferAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDITransfer(dAHDIChannel string, opts ...RequestOption) (res *Response, err error) {
	req := &DAHDITransferAction{
		DAHDIChannel: dAHDIChannel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
