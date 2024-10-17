package calculator

import "fmt"

// MainWindow Receiver получатель команды
type MainWindow struct {
	//widget *Widget
}

func NewMainWindow() *MainWindow {
	return &MainWindow{}
}

func (w *MainWindow) openBasicCalculator() {
	fmt.Println("OpenBasicCalculator")
}

func (w *MainWindow) openScientificCalculator() {
	fmt.Println("OpenScientificCalculator")
}

func (w *MainWindow) openProgrammerCalculator() {
	fmt.Println("OpenProgrammerCalculator")
}

func (w *MainWindow) MakeSwitcherBasic() *SwitchToBasic {
	return &SwitchToBasic{mw: w}
}

func (w *MainWindow) MakeSwitcherScientific() *SwitchToScientific {
	return &SwitchToScientific{mw: w}
}

func (w *MainWindow) MakeSwitcherProgrammer() *SwitchToProgrammer {
	return &SwitchToProgrammer{mw: w}
}
