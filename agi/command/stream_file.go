package command

type StreamFileCommand struct {
	FileName     string
	EscapeDigits string
	SampleOffset *int
}

func (cmd StreamFileCommand) CommandString() string {
	return joinCommand([]interface{}{"STREAM FILE", cmd.FileName, cmd.EscapeDigits, cmd.SampleOffset})
}

func (cmd StreamFileCommand) SetSampleOffset(v int) Command {
	cmd.SampleOffset = &v
	return cmd
}

func NewStreamFileCommand(fileName, escapeDigits string) Command {
	return StreamFileCommand{FileName: fileName, EscapeDigits: escapeDigits}
}
