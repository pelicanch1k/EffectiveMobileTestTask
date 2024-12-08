package main

import (
	"context"
	"github.com/joho/godotenv"
	song "github.com/pelicanch1k/EffectiveMobileTestTask"
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/handler"
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/repository"
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/repository/postgres"
	"github.com/pelicanch1k/EffectiveMobileTestTask/internal/service"
	"github.com/pelicanch1k/EffectiveMobileTestTask/pkg/logging"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// TODO: init logger
	logger := logging.GetLogger()

	// TODO: init .env
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	// TODO: init repository
	db, err := postgres.NewPostgresDB()
	if err != nil {
		logger.Fatalf("error connecting to database: %s", err.Error())
	}

	//TODO: start server
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, logger)

	srv := new(song.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Print("Server is running")

	//TODO: Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
