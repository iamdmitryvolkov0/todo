package main

import (
	"github.com/spf13/viper"
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	err := srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("../cfg")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
