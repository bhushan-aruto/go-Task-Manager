package usecase

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bhushan-aruto/go-task-manager/internal/entity"
	"github.com/bhushan-aruto/go-task-manager/internal/repository"
)

type TaskUsecase struct {
	database repository.TaskRepository
}

func NewTaskUseCaseRepo(db repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		database: db,
	}
}

func (u *TaskUsecase) CreateTaskUsecase(userId, title, description string) error {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)

	if userId == "" || title == "" || description == "" {
		return errors.New("all fields are required")
	}

	task := entity.NewTask(userId, title, description)

	if err := u.database.Creat(task); err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil

}

func (u *TaskUsecase) GetTaskByID(id string) (*entity.Task, error) {
	if id == "" {
		return nil, errors.New("task ID is required")
	}

	task, err := u.database.GetTaskById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return task, nil
}

func (u *TaskUsecase) UpdateTask(task *entity.Task) error {
	if task == nil {
		return errors.New("task cannot be nil")
	}
	if task.ID.IsZero() {
		return errors.New("invalid task ID")
	}
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(task.Description) == "" {
		return errors.New("description is required")
	}

	task.UpdatedAt = time.Now().UTC()

	if err := u.database.UpdateTask(task); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (u *TaskUsecase) DeleteTaskByID(id string) error {
	if id == "" {
		return errors.New("task ID is required")
	}

	if err := u.database.DeleteTaskById(id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}

func (u *TaskUsecase) ListTasksByUser(userId string) ([]*entity.Task, error) {
	if userId == "" {
		return nil, errors.New("user ID is required")
	}

	tasks, err := u.database.ListByUser(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	return tasks, nil
}

func (u *TaskUsecase) MarkTaskAsComplete(id string) error {
	if id == "" {
		return errors.New("task ID is required")
	}

	task, err := u.database.GetTaskById(id)
	if err != nil {
		return fmt.Errorf("task not found: %w", err)
	}

	task.Completed = true
	task.UpdatedAt = time.Now()

	if err := u.database.UpdateTask(task); err != nil {
		return fmt.Errorf("failed to mark task as complete: %w", err)
	}

	return nil
}
