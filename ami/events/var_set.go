package events

type VarSetEvent struct {
	Variable string
	Value    string
}

func (VarSetEvent) EventTypeName() string {
	return "VarSet"
}
