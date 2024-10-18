package main

import (
	conf2 "builder_pattern/internal/config"
	controller2 "builder_pattern/internal/controller"
	"fmt"
)

func main() {
	controller := controller2.NewController()
	config := conf2.NewConfig()
	controller.LoadObjectSettings(&config)
	setObject := controller.GetObjectSettings()
	fmt.Println(setObject)
}
