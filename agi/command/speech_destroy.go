package command

type SpeechDestroyCommand struct{}

func (cmd SpeechDestroyCommand) CommandString() string {
	return "SPEECH DESTROY"
}

func NewSpeechDestroyCommand() Command {
	return SpeechDestroyCommand{}
}
