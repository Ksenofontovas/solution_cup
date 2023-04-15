package scheduler

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ChatID     int64
	StartAt    time.Time
	Duration   time.Duration
	FinishAt   time.Time
	Zones      int
	IsManual   bool
	IsCritical bool
}
