package actions

type UpdateConfigAction struct {
	ActionID                 string
	SrcFileName              string
	DstFileName              string
	Reload                   string
	PreserveEffectiveContext string
	Action000000             string
	Cat000000                string
	Var000000                string
	Value000000              string
	Match000000              string
	Line000000               string
	Options000000            string
}

func (UpdateConfigAction) ActionTypeName() string {
	return "UpdateConfig"
}
func (a UpdateConfigAction) GetActionID() string {
	return a.ActionID
}
func (a *UpdateConfigAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) UpdateConfig(srcFileName string, dstFileName string, opts ...RequestOption) (res *Response, err error) {
	req := &UpdateConfigAction{
		SrcFileName: srcFileName,
		DstFileName: dstFileName,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
