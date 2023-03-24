package actions

type DAHDIDNDoffAction struct {
	ActionID     string
	DAHDIChannel string
}

func (DAHDIDNDoffAction) ActionTypeName() string {
	return "DAHDIDNDoff"
}
func (a DAHDIDNDoffAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDNDoffAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIDNDoff(dAHDIChannel string, opts ...RequestOption) (res *Response, err error) {
	req := &DAHDIDNDoffAction{
		DAHDIChannel: dAHDIChannel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
