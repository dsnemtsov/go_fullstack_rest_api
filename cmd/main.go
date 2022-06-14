package main

import (
	"github.com/spf13/viper"
	todo "go_fullstack_crud"
	"go_fullstack_crud/pkg/handler"
	"go_fullstack_crud/pkg/repository"
	"go_fullstack_crud/pkg/service"
	"log"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
