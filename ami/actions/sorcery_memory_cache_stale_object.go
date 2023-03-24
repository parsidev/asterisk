package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) SorceryMemoryCacheStaleObject(cache string, object string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SorceryMemoryCacheStaleObjectAction{
		Cache:  cache,
		Object: object,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
