package repository

import (
	"fmt"
	"github.com/dora1998/snail-bot/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type DBRepository struct {
	db   *sqlx.DB
	uuid uuid.UUID
}

func NewDBRepository(db *sqlx.DB) TaskRepository {
	repo := &DBRepository{db: db}

	u, _ := uuid.NewRandom()
	repo.uuid = u

	return repo
}

func (r *DBRepository) generateId() string {
	return r.uuid.String()
}

func (r *DBRepository) Add(body string, deadline time.Time, createdBy string) *models.Task {
	task := &models.Task{Id: r.generateId(), Body: body, Deadline: deadline, CreatedBy: createdBy}
	_, err := r.db.NamedExec("INSERT INTO tasks (id, body, deadline, created_by) VALUES (:id, :body, :deadline, :created_by)", task)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	err = r.db.Get(task, "SELECT * FROM tasks WHERE id=?", task.Id)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	return task
}

func (r *DBRepository) Remove(id string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}

func (r *DBRepository) GetAllTasks() []models.Task {
	var tasks []models.Task
	r.db.Select(&tasks, "SELECT * FROM tasks WHERE deadline > NOW() ORDER BY deadline ASC")
	return tasks
}

func (r *DBRepository) GetTaskById(id string) *models.Task {
	task := &models.Task{}
	err := r.db.Get(&task, "SELECT * FROM tasks WHERE id=$1", id)
	if err != nil {
		return nil
	}
	return task
}

func (r *DBRepository) GetTaskByBody(body string) *models.Task {
	task := &models.Task{}
	err := r.db.Get(&task, "SELECT * FROM tasks WHERE body=$1", body)
	if err != nil {
		return nil
	}
	return task
}
