package test_exactTime

import (
	"fmt"
	"github.com/beevik/ntp"
	_ "github.com/beevik/ntp"
	"os"
	"time"
)

// host определяет NTP-сервер для получения точного времени
var host string = "0.beevik-ntp.pool.ntp.org"

// GetTime функция получает точное время с NTP-сервера и выводит его в консоль
// Если возникает ошибка она выводится в Stderr, а функция возвращает код ошибки
func GetTime(params ...string) int {
	if len(params) != 0 {
	}
	// получение времени с ntp-сервера
	ntpTime, err := ntp.Time(host)
	if err != nil {
		strErr := fmt.Sprintf("Ошибка получения времени: %s", err)
		fmt.Fprint(os.Stderr, strErr)
		return -1
	}

	// Форматируем полученное NTP-время
	ntpTimeFormatted := ntpTime.Format(time.UnixDate)
	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	fmt.Println("#########################")
	// Получаем и форматируем текущее системное время
	timeFormatted := time.Now().Local().Format(time.UnixDate)
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)

	return 0
}
