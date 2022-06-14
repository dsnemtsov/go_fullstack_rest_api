package service

import "go_fullstack_crud/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos: repos}
}
