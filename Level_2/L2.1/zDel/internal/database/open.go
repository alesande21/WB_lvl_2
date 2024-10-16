package database

import (
	"facade_pattern/internal/config"
	"fmt"
)

type ConnDb struct {
	conn string
}

func Open(cfg *config.Config) *ConnDb {
	fmt.Println(cfg.GetSettings())
	return &ConnDb{conn: "Соединение с бд установлено..."}
}
