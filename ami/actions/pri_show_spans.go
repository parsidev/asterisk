package actions

type PRIShowSpansAction struct {
	ActionID string
	Span     string
}

func (PRIShowSpansAction) ActionTypeName() string {
	return "PRIShowSpans"
}
func (a PRIShowSpansAction) GetActionID() string {
	return a.ActionID
}
func (a *PRIShowSpansAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) PRIShowSpans(opts ...RequestOption) (res *Response, err error) {
	req := &PRIShowSpansAction{}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
