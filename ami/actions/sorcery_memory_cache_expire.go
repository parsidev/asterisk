package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) SorceryMemoryCacheExpire(cache string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SorceryMemoryCacheExpireAction{
		Cache: cache,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
