package command

type SpeechSetCommand struct {
	Name  string
	Value string
}

func (cmd SpeechSetCommand) CommandString() string {
	return joinCommand([]interface{}{"SPEECH SET", cmd.Name, cmd.Value})
}

func NewSpeechSetCommand(name, value string) Command {
	return SpeechSetCommand{Name: name, Value: value}
}
