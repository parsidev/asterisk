package actions

type DAHDIDialOffhookAction struct {
	ActionID     string
	DAHDIChannel string
	Number       string
}

func (DAHDIDialOffhookAction) ActionTypeName() string {
	return "DAHDIDialOffhook"
}
func (a DAHDIDialOffhookAction) GetActionID() string {
	return a.ActionID
}
func (a *DAHDIDialOffhookAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) DAHDIDialOffhook(dAHDIChannel string, number string, opts ...RequestOption) (res *Response, err error) {
	req := &DAHDIDialOffhookAction{
		DAHDIChannel: dAHDIChannel,
		Number:       number,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
