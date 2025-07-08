package repository

import "github.com/bhushan-aruto/go-task-manager/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
