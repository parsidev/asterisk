package command

type AsyncAgiBreakCommand struct{}

func (cmd AsyncAgiBreakCommand) CommandString() string {
	return "ASYNCAGI BREAK"
}

func NewAsyncAgiBreakCommand() Command {
	return AsyncAgiBreakCommand{}
}
