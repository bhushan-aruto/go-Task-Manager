package repository

import "github.com/bhushan-aruto/go-task-manager/internal/entity"

type TaskRepository interface {
	Creat(task *entity.Task) error
	GetTaskById(id string) (*entity.Task, error)
	UpdateTask(task *entity.Task) error
	DeleteTaskById(id string) error
}
