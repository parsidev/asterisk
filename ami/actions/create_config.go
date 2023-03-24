package actions

import "github.com/parsidev/asterisk/ami/message"

type CreateConfigAction struct {
	ActionID string
	FileName string
}

func (CreateConfigAction) ActionTypeName() string {
	return "CreateConfig"
}
func (a CreateConfigAction) GetActionID() string {
	return a.ActionID
}
func (a *CreateConfigAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) CreateConfig(fileName string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &CreateConfigAction{
		FileName: fileName,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
