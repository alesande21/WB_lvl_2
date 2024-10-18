package main

import "factory_pattern/internal/view"

func main() {
	mw := view.NewMainWindow()
	mw.OpenTempWindow("color")
	mw.OpenTempWindow("texture")
	mw.OpenTempWindow("camera")
}

/*
Отркыты настройки цвета.
Отркыты настройки текстуры.
Отркыты настройки камеры.
*/
