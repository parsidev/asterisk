package command

type ControlStreamFileCommand struct {
	FileName     string
	EscapeDigits string
	SkipMS       *int
	FfChar       *string
	RewChr       *string
	PauseChr     *string
	OffsetMS     *int
}

func (cmd ControlStreamFileCommand) CommandString() string {
	return joinCommand([]interface{}{"CONTROL STREAM FILE", cmd.FileName, cmd.EscapeDigits, cmd.SkipMS,
		cmd.FfChar, cmd.RewChr, cmd.PauseChr, cmd.OffsetMS})
}

func (cmd ControlStreamFileCommand) SetSkipMS(v int) Command {
	cmd.SkipMS = &v
	return cmd
}
func (cmd ControlStreamFileCommand) SetFfChar(v string) Command {
	cmd.FfChar = &v
	return cmd
}
func (cmd ControlStreamFileCommand) SetRewChr(v string) Command {
	cmd.RewChr = &v
	return cmd
}
func (cmd ControlStreamFileCommand) SetPauseChr(v string) Command {
	cmd.PauseChr = &v
	return cmd
}
func (cmd ControlStreamFileCommand) SetOffsetMS(v int) Command {
	cmd.OffsetMS = &v
	return cmd
}

func NewControlStreamFileCommand(fileName, escapeDigits string) Command {
	return ControlStreamFileCommand{
		FileName:     fileName,
		EscapeDigits: escapeDigits,
	}
}
