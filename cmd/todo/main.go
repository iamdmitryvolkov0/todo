package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"todo"
	"todo/internal/handler"
	"todo/internal/repository"
	"todo/internal/service"
	"todo/internal/storage"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %s", err)
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err)
	}

	db, err := storage.NewPostgresDB(storage.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Error initializing database: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error starting server: %s", err.Error())
		}
	}()

	logrus.Info("Application started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	logrus.Info("Application shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Error while shutting down server: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Fatalf("Error while closing db connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
