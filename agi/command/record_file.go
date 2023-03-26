package command

type RecordFileCommand struct {
	FileName      string
	Format        string
	EscapeDigits  string
	Timeout       int
	OffsetSamples *string
	Beep          *string
	SSilence      *string
}

func (cmd RecordFileCommand) CommandString() string {
	return joinCommand([]interface{}{"RECORD FILE", cmd.FileName, cmd.Format, cmd.EscapeDigits, cmd.Timeout,
		cmd.OffsetSamples, cmd.Beep, cmd.SSilence})
}

func (cmd RecordFileCommand) SetOffsetSamples(v string) Command {
	cmd.OffsetSamples = &v
	return cmd
}
func (cmd RecordFileCommand) SetBeep(v string) Command {
	cmd.Beep = &v
	return cmd
}
func (cmd RecordFileCommand) SetSSilence(v string) Command {
	cmd.SSilence = &v
	return cmd
}

func NewRecordFileCommand(fileName, format, escapeDigits string, timeout int) Command {
	return RecordFileCommand{FileName: fileName, Format: format, EscapeDigits: escapeDigits, Timeout: timeout}
}
