package actions

type SorceryMemoryCacheExpireObjectAction struct {
	ActionID string
	Cache    string
	Object   string
}

func (SorceryMemoryCacheExpireObjectAction) ActionTypeName() string {
	return "SorceryMemoryCacheExpireObject"
}
func (a SorceryMemoryCacheExpireObjectAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCacheExpireObjectAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SorceryMemoryCacheExpireObject(cache string, object string, opts ...RequestOption) (res *Response, err error) {
	req := &SorceryMemoryCacheExpireObjectAction{
		Cache:  cache,
		Object: object,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
