package repository

import (
	"github.com/dora1998/snail-bot/models"
	"time"
)

type TaskRepository interface {
	Add(body string, deadline time.Time, createdBy string) *models.Task
	Remove(id string) error
	GetAllTasks() []models.Task
	GetTaskById(id string) *models.Task
	GetTaskByBody(name string) *models.Task
}
