package app

func Run() error {
	// Настройка логера
	SetLevel("debug", "console")
	log2.Info("Настройка логера...")

	// Инициализация базы данных
	log.Println("Инициализация базы данных...")
	var conn *database2.DBConnection
	conn, err = database2.Open(config.GetDBsConfig())
	if err != nil {
		return fmt.Errorf("-> database2.Open%w", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log2.Infof("RunConsumer-> conn.Close:%s", err)
		}
	}()
}
