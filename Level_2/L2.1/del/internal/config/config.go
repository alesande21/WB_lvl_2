package config

import "fmt"

type Config struct {
	settings string
}

func GetConfig() *Config {
	fmt.Println("Конфиг получен...")
	cnfg := Config{settings: "Конфиг настроен..."}
	return &cnfg
}

func (c Config) GetSettings() string {
	return c.settings
}
