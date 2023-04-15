package service

import (
	"github.com/Ksenofontovas/solution_cup/internal/repository"
)

type Task interface {
	ValidateTask(message string) bool
}

type Service struct {
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Task: NewTaskService(repos),
	}
}
