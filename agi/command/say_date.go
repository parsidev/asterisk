package command

type SayDateCommand struct {
	Date         string
	EscapeDigits string
}

func (cmd SayDateCommand) CommandString() string {
	return joinCommand([]interface{}{"SAY DATE", cmd.Date, cmd.EscapeDigits})
}

func NewSayDateCommand(date, escapeDigits string) Command {
	return SayDateCommand{Date: date, EscapeDigits: escapeDigits}
}
