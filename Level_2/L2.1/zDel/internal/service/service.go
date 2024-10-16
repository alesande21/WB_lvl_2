package service

import (
	"facade_pattern/internal/database"
	"fmt"
)

type Service struct {
	command string
}

func GetService(repository *database.Repository) *Service {
	fmt.Println(repository.Get())
	return &Service{command: "Приложение запущено..."}
}

func (s Service) GetCommand() string {
	return s.command
}
