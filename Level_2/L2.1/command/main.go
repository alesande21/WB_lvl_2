package main

import "command_pattern/internal/calculator"

func main() {
	mw := calculator.NewMainWindow()
	invoker := calculator.NewInvoker()
	invoker.SetCommand(mw.MakeSwitcherBasic()).
		SetCommand(mw.MakeSwitcherProgrammer()).
		SetCommand(mw.MakeSwitcherScientific()).
		ExecuteCommands()
}
