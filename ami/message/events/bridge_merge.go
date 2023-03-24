package events

type BridgeMergeEvent struct {
}

func (BridgeMergeEvent) EventTypeName() string {
	return "BridgeMerge"
}
