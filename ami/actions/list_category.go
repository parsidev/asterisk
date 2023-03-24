package actions

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
func (cli *Client) ListCategories(fileName string, opts ...RequestOption) (res *Response, err error) {
	req := &ListCategoriesAction{
		FileName: fileName,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
