package main

import (
	"fmt"
	"github.com/google/uuid"
)

type TaskRepository struct {
	tasks      []*Task
	uuid       uuid.UUID
	generateId func() string
}

func NewTaskRepository() (*TaskRepository, error) {
	repository := &TaskRepository{tasks: make([]*Task, 0)}

	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	repository.uuid = u

	repository.generateId = repository.uuid.String

	return repository, nil
}

func (r *TaskRepository) Add(body string, createdBy string) *Task {
	task := &Task{Id: r.generateId(), Body: body, CreatedBy: createdBy}
	fmt.Printf("%#v\n", task)
	r.tasks = append(r.tasks, task)
	return task
}

func (r *TaskRepository) Remove(id string) {
	res := make([]*Task, 0)
	for _, task := range r.tasks {
		if task.Id != id {
			res = append(res, task)
		}
	}
	r.tasks = res
}

func (r *TaskRepository) GetAllTasks() []*Task {
	return r.tasks
}

func (r *TaskRepository) GetTaskById(id string) *Task {
	for _, task := range r.tasks {
		if task.Id != id {
			return task
		}
	}
	return nil
}
