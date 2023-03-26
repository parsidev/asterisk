package command

type SpeechLoadGrammarCommand struct {
	GrammarName   string
	PathToGrammar string
}

func (cmd SpeechLoadGrammarCommand) CommandString() string {
	return joinCommand([]interface{}{"SPEECH LOAD GRAMMAR", cmd.GrammarName, cmd.PathToGrammar})
}

func NewSpeechLoadGrammarCommand(name, path string) Command {
	return SpeechLoadGrammarCommand{GrammarName: name, PathToGrammar: path}
}
