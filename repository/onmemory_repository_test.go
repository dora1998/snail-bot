package repository

import (
	"github.com/dora1998/snail-bot/models"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

func TestTaskRepository_Add(t *testing.T) {
	type fields struct {
		tasks      []models.Task
		uuid       uuid.UUID
		generateId func() string
	}
	type args struct {
		body      string
		deadline  time.Time
		createdBy string
	}
	type want struct {
		got   *models.Task
		tasks []models.Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "test",
			fields: fields{
				tasks: make([]models.Task, 0),
				uuid:  uuid.UUID{},
				generateId: func() string {
					return "hoge"
				},
			},
			args: args{
				body:      "task test",
				deadline:  time.Time{},
				createdBy: "2019",
			},
			want: want{
				got: &models.Task{
					Id:        "hoge",
					Body:      "task test",
					Deadline:  time.Time{},
					CreatedAt: time.Time{},
					CreatedBy: "2019",
				},
				tasks: []models.Task{{
					Id:        "hoge",
					Body:      "task test",
					Deadline:  time.Time{},
					CreatedAt: time.Time{},
					CreatedBy: "2019",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := OnMemoryRepository{
				tasks:      tt.fields.tasks,
				uuid:       tt.fields.uuid,
				generateId: tt.fields.generateId,
			}
			got := r.Add(tt.args.body, tt.args.deadline, tt.args.createdBy)
			if !reflect.DeepEqual(got, tt.want.got) {
				t.Errorf("Add() = %#v, want %#v", got, tt.want.got)
			}
			if !reflect.DeepEqual(r.tasks, tt.want.tasks) {
				t.Errorf("Add() = %#v, want %#v", r.tasks, tt.want.tasks)
			}
		})
	}
}

func TestTaskRepository_Remove(t *testing.T) {
	type fields struct {
		tasks      []models.Task
		uuid       uuid.UUID
		generateId func() string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.Task
	}{
		{
			name: "test",
			fields: fields{
				tasks: []models.Task{{
					Id:        "hoge",
					Body:      "test body",
					Deadline:  time.Time{},
					CreatedAt: time.Time{},
					CreatedBy: "d0ra1998",
				}},
				uuid:       uuid.UUID{},
				generateId: nil,
			},
			args: args{
				id: "hoge",
			},
			want: []models.Task{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := OnMemoryRepository{
				tasks:      tt.fields.tasks,
				uuid:       tt.fields.uuid,
				generateId: tt.fields.generateId,
			}
			if r.Remove(tt.args.id); !reflect.DeepEqual(r.tasks, tt.want) {
				t.Errorf("Remove() = %#v, want %#v", r.tasks, tt.want)
			}
		})
	}
}