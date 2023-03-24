package actions

type SorceryMemoryCacheStaleAction struct {
	ActionID string
	Cache    string
}

func (SorceryMemoryCacheStaleAction) ActionTypeName() string {
	return "SorceryMemoryCacheStale"
}
func (a SorceryMemoryCacheStaleAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCacheStaleAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SorceryMemoryCacheStale(cache string, opts ...RequestOption) (res *Response, err error) {
	req := &SorceryMemoryCacheStaleAction{
		Cache: cache,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
