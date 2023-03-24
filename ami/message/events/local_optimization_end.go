package events

type LocalOptimizationEndEvent struct {
	Success string
	Id      string
}

func (LocalOptimizationEndEvent) EventTypeName() string {
	return "LocalOptimizationEnd"
}
