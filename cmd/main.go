package main

import (
	todo "go_fullstack_crud"
	"go_fullstack_crud/pkg/handler"
	"go_fullstack_crud/pkg/repository"
	"go_fullstack_crud/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err.Error())
	}
}
