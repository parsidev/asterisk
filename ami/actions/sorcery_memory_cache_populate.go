package actions

type SorceryMemoryCachePopulateAction struct {
	ActionID string
	Cache    string
}

func (SorceryMemoryCachePopulateAction) ActionTypeName() string {
	return "SorceryMemoryCachePopulate"
}
func (a SorceryMemoryCachePopulateAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCachePopulateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SorceryMemoryCachePopulate(cache string, opts ...RequestOption) (res *Response, err error) {
	req := &SorceryMemoryCachePopulateAction{
		Cache: cache,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
