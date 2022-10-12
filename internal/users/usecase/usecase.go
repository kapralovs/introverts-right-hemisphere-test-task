package usecase

import (
	"github.com/kapralovs/simple-test-api/internal/models"
	"github.com/kapralovs/simple-test-api/internal/users"
)

type UserUsecase struct {
	repo users.Repository
}

func New(r users.Repository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (uc *UserUsecase) GetUsers() ([]*models.User, error) {
	users, err := uc.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUsecase) EditUser(data *models.User) error {
	err := uc.repo.EditUser(data)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) DeleteUser(id string) error {
	err := uc.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
