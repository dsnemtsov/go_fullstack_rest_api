package main

import (
	todo "go_fullstack_crud"
	"go_fullstack_crud/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err.Error())
	}
}
