package actions

type SorceryMemoryCacheExpireAction struct {
	ActionID string
	Cache    string
}

func (SorceryMemoryCacheExpireAction) ActionTypeName() string {
	return "SorceryMemoryCacheExpire"
}
func (a SorceryMemoryCacheExpireAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCacheExpireAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SorceryMemoryCacheExpire(cache string, opts ...RequestOption) (res *Response, err error) {
	req := &SorceryMemoryCacheExpireAction{
		Cache: cache,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
