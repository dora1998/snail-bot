package main

import "time"

type Task struct {
	Id        string    `db:"id"`
	Body      string    `db:"body"`
	Deadline  time.Time `db:"deadline"`
	CreatedAt time.Time `db:"created_at"`
	CreatedBy string    `db:"created_by"`
}
