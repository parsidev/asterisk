package command

type SpeechUnloadGrammarCommand struct {
	GrammarName string
}

func (cmd SpeechUnloadGrammarCommand) CommandString() string {
	return joinCommand([]interface{}{"SPEECH UNLOAD GRAMMAR", cmd.GrammarName})
}

func NewSpeechUnloadGrammarCommand(name string) Command {
	return SpeechUnloadGrammarCommand{GrammarName: name}
}
