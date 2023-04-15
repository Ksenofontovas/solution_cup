package service

import (
	"strings"

	scheduler "github.com/Ksenofontovas/solution_cup/domain"
	"github.com/Ksenofontovas/solution_cup/internal/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task scheduler.Task) (uint, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) ValidateTask(message string) bool {
	return strings.Contains(message, "Новая работа")
}
