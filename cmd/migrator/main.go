package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/pelicanch1k/EffectiveMobileTestTask/pkg/logging"
	"os"
)

func main() {
	logger := logging.GetLogger()

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	migrationsPath := os.Getenv("MIGRATIONS_PATH")

	// Создаём экземпляр миграций
	m, err := migrate.New(migrationsPath, databaseURL())
	if err != nil {
		logger.Fatalf("Ошибка создания миграции: %v", err)
	}

	// Применяем миграции "up"
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Println("Миграции не применялись: изменений нет.")
		} else {
			logger.Fatalf("Ошибка применения миграций: %v", err)
		}
	} else {
		logger.Println("Миграции успешно применены!")
	}
}

func databaseURL() string {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	sslmode := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, database, sslmode)
}
