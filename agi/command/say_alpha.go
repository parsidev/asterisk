package command

type SayAlphaCommand struct {
	Number       string
	EscapeDigits string
}

func (cmd SayAlphaCommand) CommandString() string {
	return joinCommand([]interface{}{"SAY ALPHA", cmd.Number, cmd.EscapeDigits})
}

func NewSayAlphaCommand(number, escapeDigits string) Command {
	return SayAlphaCommand{Number: number, EscapeDigits: escapeDigits}
}
