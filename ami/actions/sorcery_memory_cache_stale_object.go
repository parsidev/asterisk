package actions

type SorceryMemoryCacheStaleObjectAction struct {
	ActionID string
	Cache    string
	Object   string
	Reload   string
}

func (SorceryMemoryCacheStaleObjectAction) ActionTypeName() string {
	return "SorceryMemoryCacheStaleObject"
}
func (a SorceryMemoryCacheStaleObjectAction) GetActionID() string {
	return a.ActionID
}
func (a *SorceryMemoryCacheStaleObjectAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SorceryMemoryCacheStaleObject(cache string, object string, opts ...RequestOption) (res *Response, err error) {
	req := &SorceryMemoryCacheStaleObjectAction{
		Cache:  cache,
		Object: object,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
