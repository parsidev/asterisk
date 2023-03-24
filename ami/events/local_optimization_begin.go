package events

type LocalOptimizationBeginEvent struct {
	DestUniqueId string
	Id           string
}

func (LocalOptimizationBeginEvent) EventTypeName() string {
	return "LocalOptimizationBegin"
}
