package app

import (
	"calendarEvent/internal/colorAttribute"
	config2 "calendarEvent/internal/config"
	http2 "calendarEvent/internal/controllers/http"
	database2 "calendarEvent/internal/database"
	repository2 "calendarEvent/internal/repository"
	service2 "calendarEvent/internal/service"
	"context"
	"errors"
	"fmt"
	log2 "github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	//Загрузка конфига
	log.Println("Загрузка конфига...")
	config, err := config2.GetDefaultConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Проблема с загрузкой конфига\n: %s", err)
		return nil
	}

	// Инициализация базы данных
	log.Println("Инициализация базы данных...")
	var conn *database2.DBConnection
	conn, err = database2.Open(config.GetDBsConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "проблемы с драйвером подключения на этапе открытия\n: %s", err)
		return nil

	}
	defer func() {
		if err := conn.Close(); err != nil {
			log2.Infof("RunConsumer-> conn.Close:%s", err)
		}
	}()

	// Инициализация репозитория
	log.Println("Инициализация репозитория...")
	var postgresRep database2.DBRepository
	postgresRep, err = database2.CreatePostgresRepository(conn.GetConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "проблемы с созданием PostgresRepository \n: %s", err)
		return nil

	}

	//err = InitSchema(postgresRep, "migrations/schema.sql")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "не удалось загрузить схему \n: %s", err)
	//	return
	//}

	log.Println("Инициализация сервиса...")
	eventRepo := repository2.NewEventRepo(postgresRep)
	eventService := service2.NewOrderService(eventRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Загрузка настроек для сервера...")
	var serverAddress http2.ServerAddress
	err = serverAddress.UpdateEnvAddress()
	if err != nil {
		fmt.Fprintf(os.Stderr, "настройки адреса сервера не загрузились\n: %s", err)
		return nil

	}

	log.Println("Инициализация и старт сервера...")
	//swagger, err := http2.GetSwagger()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "ошибка загрузки сваггера\n: %s", err)
	//	return nil
	//
	//}
	//swagger.Servers = nil

	r := http2.NewRouter(eventService)

	s := &http.Server{
		Addr:    serverAddress.EnvAddress,
		Handler: r,
	}
	//tenderServer := http2.NewTenderServer(orderService)

	//r := mux.NewRouter()

	//r.Use(middleware.OapiRequestValidator(swagger))
	//http2.HandlerFromMux(tenderServer, r)

	//s := &http.Server{
	//	Addr:    serverAddress.EnvAddress,
	//	Handler: r,
	//}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer close(interrupt)

	shutDownChan := make(chan error, 1)
	defer close(shutDownChan)

	// Проверка подключения
	go func() {
		shutDownChan <- conn.CheckConn(ctx, config.GetDBsConfig())
	}()

	go func() {
		shutDownChan <- s.ListenAndServe()
	}()

	log2.Infof("Подключнеие установлено -> %s", colorAttribute.ColorString(colorAttribute.FgYellow, serverAddress.EnvAddress))

	select {
	case sig := <-interrupt:
		log2.Infof("Приложение прерывается: %s", sig)
		ctxShutDown, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)

		cancel()

		defer cancelShutdown()
		err := s.Shutdown(ctxShutDown)
		if err != nil {
			return fmt.Errorf("-> s.Shutdown: %w", err)
		}

		log2.Info("Сервер завершил работу")
	case err := <-shutDownChan:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf(": ошибка при запуске сервера: %w", err)
		}
	}

	return nil

}
