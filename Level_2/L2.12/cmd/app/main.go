package app

import (
	"calendarEvent/internal/app"
	"fmt"
)

func main() {
	err := app.Run()
	if err != nil {
		fmt.Printf("Ошибка:%v", err)
	}
}
