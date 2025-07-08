package usecase

import (
	"errors"

	"github.com/bhushan-aruto/go-task-manager/internal/entity"
	"github.com/bhushan-aruto/go-task-manager/internal/repository"
	"github.com/bhushan-aruto/go-task-manager/internal/utils"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) Register(user *entity.User) error {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	return u.repo.Create(user)
}

func (u *UserUsecase) Login(email, password string) (*entity.User, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
