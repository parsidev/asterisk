package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) SorceryMemoryCacheExpireObject(cache string, object string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SorceryMemoryCacheExpireObjectAction{
		Cache:  cache,
		Object: object,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
