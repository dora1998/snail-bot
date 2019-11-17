package repository

import (
	"time"
)

type TaskRepository interface {
	Add(body string, deadline time.Time, createdBy string) *Task
	Remove(id string) error
	GetAllTasks() []Task
	GetTaskById(id string) *Task
	GetTaskByBody(body string) *Task
}
