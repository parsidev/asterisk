package command

type SayTimeCommand struct {
	Time         float64
	EscapeDigits string
}

func (cmd SayTimeCommand) CommandString() string {
	return joinCommand([]interface{}{"SAY TIME", cmd.Time, cmd.EscapeDigits})
}

func NewSayTimeCommand(time float64, escapeDigits string) Command {
	return SayTimeCommand{Time: time, EscapeDigits: escapeDigits}
}
