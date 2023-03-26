package command

type SpeechActivateGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechActivateGrammarCommand) CommandString() string {
	return joinCommand([]interface{}{"SPEECH ACTIVATE GRAMMAR", cmd.GrammarName})
}

func NewSpeechActivateGrammarCommand(name string) Command {
	return SpeechActivateGrammarCommand{GrammarName: name}
}
