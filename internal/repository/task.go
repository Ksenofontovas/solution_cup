package repository

import (
	scheduler "github.com/Ksenofontovas/solution_cup/domain"
	"gorm.io/gorm"
)

type TaskRepos struct {
	db *gorm.DB
}

func NewTaskRepos(db *gorm.DB) *TaskRepos {
	return &TaskRepos{db: db}
}

func (r *TaskRepos) CreateTask(task scheduler.Task) (uint, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return 0, result.Error
	}
	return task.ID, nil

}
