package actions

type GetConfigAction struct {
	ActionID string
	FileName string
	Category string
	Filter   string
}

func (GetConfigAction) ActionTypeName() string {
	return "GetConfig"
}
func (a GetConfigAction) GetActionID() string {
	return a.ActionID
}
func (a *GetConfigAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) GetConfig(fileName string, opts ...RequestOption) (res *Response, err error) {
	req := &GetConfigAction{
		FileName: fileName,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
