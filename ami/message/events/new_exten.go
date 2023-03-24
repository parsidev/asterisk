package events

type NewExtenEvent struct {
	Extension   string
	Application string
	AppData     string
}

func (NewExtenEvent) EventTypeName() string {
	return "NewExten"
}
