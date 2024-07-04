package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"todo"
	"todo/internal/handler"
	"todo/internal/repository"
	"todo/internal/service"
	"todo/internal/storage"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err)
	}

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("error loading .env file: %s", err)
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
		log.Fatalf("error initializing database: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("../cfg")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
