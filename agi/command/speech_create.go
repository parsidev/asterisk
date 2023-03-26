package command

type SpeechCreateCommand struct {
	Engine string
}

func (cmd SpeechCreateCommand) CommandString() string {
	return joinCommand([]interface{}{"SPEECH CREATE", cmd.Engine})
}

func NewSpeechCreateCommand(engine string) Command {
	return SpeechCreateCommand{Engine: engine}
}
