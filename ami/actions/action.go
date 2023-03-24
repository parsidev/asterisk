package actions

type Action interface {
	ActionTypeName() string
}

type HasActionID interface {
	GetActionID() string
	SetActionID(actionID string)
}
