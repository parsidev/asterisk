package command

type SayDigitsCommand struct {
	Number       string
	EscapeDigits string
}

func (cmd SayDigitsCommand) CommandString() string {
	return joinCommand([]interface{}{"SAY DIGITS", cmd.Number, cmd.EscapeDigits})
}

func NewSayDigitsCommand(number, escapeDigits string) Command {
	return SayDigitsCommand{Number: number, EscapeDigits: escapeDigits}
}
