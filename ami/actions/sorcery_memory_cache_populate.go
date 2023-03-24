package actions

import "github.com/parsidev/asterisk/ami/message"

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
func (cli *Client) SorceryMemoryCachePopulate(cache string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &SorceryMemoryCachePopulateAction{
		Cache: cache,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
