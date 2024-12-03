package main

import (
	"fmt"
	"log"
	"os"
	"testApp/config"

	"github.com/caarlos0/env/v9"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func run() error {
	// конфигурация из .env файла
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("unable to load .env file: %e", err)
	}

	cfg := config.Config{}

	err = env.Parse(&cfg)
	if err != nil {
		return fmt.Errorf("unable to parse env variables: %e", err)
	}

	// инициализация БД и файлов миграций
	m, err := migrate.New(
		fmt.Sprintf("file://%s", cfg.MigrationsPath),
		cfg.PostgresUrl,
	)
	if err != nil {
		return fmt.Errorf("error while create migration: %v", err)
	}

	// Применение миграций
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error while apply migration: %v", err)
	}
	infoLogger.Println("Miration succesfully applied")
	return nil
}
