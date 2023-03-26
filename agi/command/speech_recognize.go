package command

type SpeechRecognizeCommand struct {
	Prompt  string
	Timeout int
	Offset  *string
}

func (cmd SpeechRecognizeCommand) CommandString() string {
	return joinCommand([]interface{}{"SPEECH RECOGNIZE", cmd.Prompt, cmd.Timeout, cmd.Offset})
}

func (cmd SpeechRecognizeCommand) SetOffset(v string) Command {
	cmd.Offset = &v
	return cmd
}

func NewSpeechRecognizeCommand(prompt string, timeout int) Command {
	return SpeechRecognizeCommand{Prompt: prompt, Timeout: timeout}
}
