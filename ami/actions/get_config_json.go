package actions

import "github.com/parsidev/asterisk/ami/message"

type GetConfigJSONAction struct {
	ActionID string
	FileName string
	Category string
	Filter   string
}

func (GetConfigJSONAction) ActionTypeName() string {
	return "GetConfigJSON"
}
func (a GetConfigJSONAction) GetActionID() string {
	return a.ActionID
}
func (a *GetConfigJSONAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) GetConfigJSON(fileName string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &GetConfigJSONAction{
		FileName: fileName,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
