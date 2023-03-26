package command

type SpeechDeactivateGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechDeactivateGrammarCommand) CommandString() string {
	return joinCommand([]interface{}{"SPEECH DEACTIVATE GRAMMAR", cmd.GrammarName})
}

func NewSpeechDeactivateGrammarCommand(name string) Command {
	return SpeechDeactivateGrammarCommand{GrammarName: name}
}
