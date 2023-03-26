package command

type AnswerCommand struct{}

func (cmd AnswerCommand) CommandString() string {
	return "ANSWER"
}

func NewAnswerCommand() Command {
	return AnswerCommand{}
}
