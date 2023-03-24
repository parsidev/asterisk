package actions

import "github.com/parsidev/asterisk/ami/message"

type ListCategoriesAction struct {
	ActionID string
	FileName string
}

func (ListCategoriesAction) ActionTypeName() string {
	return "ListCategories"
}
func (a ListCategoriesAction) GetActionID() string {
	return a.ActionID
}
func (a *ListCategoriesAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) ListCategories(fileName string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &ListCategoriesAction{
		FileName: fileName,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
