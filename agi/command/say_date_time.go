package command

type SayDatetimeCommand struct {
	Time         float64
	EscapeDigits string
	Format       *string
	Timezone     *string
}

func (cmd SayDatetimeCommand) CommandString() string {
	return joinCommand([]interface{}{"SAY DATETIME", cmd.Time, cmd.EscapeDigits, cmd.Format, cmd.Timezone})
}

func (cmd SayDatetimeCommand) SetFormat(v string) Command {
	cmd.Format = &v
	return cmd
}
func (cmd SayDatetimeCommand) SetTimezone(v string) Command {
	cmd.Timezone = &v
	return cmd
}

func NewSayDateTimeCommand(time float64, escapeDigits string) Command {
	return SayDatetimeCommand{Time: time, EscapeDigits: escapeDigits}
}
