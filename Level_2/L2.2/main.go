package main

import (
	exactTimeV2 "github.com/alesande21/exactTime/v2"
	"os"
)

// функция запускает программу, которая вызывает функцию GetTime для получения и отображения точного времени
func main() {
	os.Exit(exactTimeV2.GetTime())
}
