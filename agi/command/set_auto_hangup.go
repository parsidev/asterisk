package command

type SetAutoHangupCommand struct {
	Time float64
}

func (cmd SetAutoHangupCommand) CommandString() string {
	return joinCommand([]interface{}{"SET AUTOHANGUP", cmd.Time})
}

func NewSetAutoHangupCommand(time float64) Command {
	return SetAutoHangupCommand{Time: time}
}
