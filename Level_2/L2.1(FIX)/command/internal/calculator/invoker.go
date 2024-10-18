package calculator

// SwitchToBasic, SwitchToScientific, SwitchToProgrammer
type Command interface {
	Execute()
}

type Invoker struct {
	commands []Command
}

func NewInvoker() *Invoker {
	return &Invoker{commands: make([]Command, 0)}
}

func (i *Invoker) SetCommand(command Command) *Invoker {
	i.commands = append(i.commands, command)
	return i
}

func (i *Invoker) ExecuteCommands() {
	for _, c := range i.commands {
		c.Execute()
	}
}
