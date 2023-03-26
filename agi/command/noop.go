package command

type NoopCommand struct {
}

func (cmd NoopCommand) CommandString() string {
	return "NOOP"
}

func NewNoopCommand() Command {
	return NoopCommand{}
}
