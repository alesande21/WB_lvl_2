package calculator

type SwitchToBasic struct {
	mw *MainWindow
}

func (b *SwitchToBasic) Execute() {
	b.mw.openBasicCalculator()
}

type SwitchToProgrammer struct {
	mw *MainWindow
}

func (b *SwitchToProgrammer) Execute() {
	b.mw.openProgrammerCalculator()
}

type SwitchToScientific struct {
	mw *MainWindow
}

func (b *SwitchToScientific) Execute() {
	b.mw.openScientificCalculator()
}
