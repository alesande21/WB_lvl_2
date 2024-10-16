package logger

import "fmt"

func SetLevel(lvl string, formatter string) {
	if lvl != "" || formatter != "" {
		fmt.Println("Логгер настроен...")
	}
}
