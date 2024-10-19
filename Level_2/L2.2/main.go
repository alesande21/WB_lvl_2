package main

import (
	"fmt"
	ntp2 "github.com/beevik/ntp"
)

func main() {
	time, err := ntp2.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time)
}
