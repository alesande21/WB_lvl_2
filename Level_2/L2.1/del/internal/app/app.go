package app

import (
	config2 "facade_pattern/internal/config"
	database2 "facade_pattern/internal/database"
	"facade_pattern/internal/logger"
	"facade_pattern/internal/service"
	"fmt"
	"time"
)

func Run() {
	logger.SetLevel("debug", "console")
	time.Sleep(time.Second)
	config := config2.GetConfig()
	time.Sleep(time.Second)
	conn := database2.Open(config)
	time.Sleep(time.Second)
	repo := database2.InitRepo(conn)
	time.Sleep(time.Second)
	serv := service.GetService(repo)
	time.Sleep(time.Second)
	fmt.Println(serv.GetCommand())
	<-time.After(time.Second * 6)
	fmt.Println("Приложение завершено...")
}
