package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) SorceryMemoryCacheStale(cache string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SorceryMemoryCacheStaleAction{
		Cache: cache,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
