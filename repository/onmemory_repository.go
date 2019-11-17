package repository

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type OnMemoryRepository struct {
	tasks      []Task
	uuid       uuid.UUID
	generateId func() string
}

func NewOnMemoryRepository() (*OnMemoryRepository, error) {
	repository := &OnMemoryRepository{tasks: make([]Task, 0)}

	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	repository.uuid = u

	repository.generateId = repository.uuid.String

	return repository, nil
}

func (r *OnMemoryRepository) Add(body string, deadline time.Time, createdBy string) *Task {
	task := Task{Id: r.generateId(), Body: body, Deadline: deadline, CreatedBy: createdBy, CreatedAt: time.Now()}
	fmt.Printf("%#v\n", task)
	r.tasks = append(r.tasks, task)
	return &task
}

func (r *OnMemoryRepository) Remove(id string) error {
	res := make([]Task, 0)
	for _, task := range r.tasks {
		if task.Id != id {
			res = append(res, task)
		}
	}
	r.tasks = res
	return nil
}

func (r *OnMemoryRepository) GetAllTasks() []Task {
	return r.tasks
}

func (r *OnMemoryRepository) GetTaskById(id string) *Task {
	for _, task := range r.tasks {
		if task.Id != id {
			return &task
		}
	}
	return nil
}

func (r *OnMemoryRepository) GetTaskByBody(body string) *Task {
	for _, task := range r.tasks {
		if task.Body != body {
			return &task
		}
	}
	return nil
}
