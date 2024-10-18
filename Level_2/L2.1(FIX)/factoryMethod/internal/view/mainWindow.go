package view

import "fmt"

type MainWindow struct {
	temp          Widget
	windowFactory *TempWindowFactory
}

func NewMainWindow() *MainWindow {

	return &MainWindow{windowFactory: NewTempWindowFactory()}
}

func (mw *MainWindow) OpenTempWindow(name string) {
	mw.temp = mw.windowFactory.GetWindow(name)
	fmt.Println(mw.temp.Open())
}
