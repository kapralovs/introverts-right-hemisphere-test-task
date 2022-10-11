package usecase

import (
	"github.com/kapralovs/introverts-right-hemisphere-test-task/internal/models"
	"github.com/kapralovs/introverts-right-hemisphere-test-task/internal/users"
)

type UserUsecase struct {
	repo users.Repository
}

func New(r users.Repository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (uc *UserUsecase) Get() ([]*models.User, error) {
	users := []*models.User{}
	return users, nil
}

func (uc *UserUsecase) Edit(id int) error {
	return nil
}

func (uc *UserUsecase) Delete(id int) error {
	return nil
}
