package main

import "time"

type Task struct {
	Id        string
	Body      string
	Deadline  time.Time
	CreatedBy string
}
