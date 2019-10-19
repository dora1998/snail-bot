package main

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type DBRepository struct{}

func (repo *DBRepository) Add(body string, deadline time.Time, createdBy string) *Task {
	panic("implement me")
}

func (repo *DBRepository) Remove(id string) {
	panic("implement me")
}

func (repo *DBRepository) GetAllTasks() []*Task {
	panic("implement me")
}

func (repo *DBRepository) GetTaskById(id string) *Task {
	panic("implement me")
}

func (repo *DBRepository) GetTaskByBody(name string) *Task {
	panic("implement me")
}

func NewDBRepository(db *sqlx.DB) TaskRepository {
	repo := &DBRepository{}
	return repo
}
