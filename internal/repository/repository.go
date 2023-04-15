package repository

import (
	scheduler "github.com/Ksenofontovas/solution_cup/domain"
	"gorm.io/gorm"
)

type Task interface {
	CreateTask(tast scheduler.Task) (uint, error)
}

type Repository struct {
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Task: NewTaskRepos(db),
	}
}
