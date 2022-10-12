package users

import "github.com/kapralovs/simple-test-api/internal/models"

type Usecase interface {
	GetUsers() ([]*models.User, error)
	EditUser(data *models.User) error
	DeleteUser(id string) error
}

type Repository interface {
	GetUsers() ([]*models.User, error)
	EditUser(data *models.User) error
	DeleteUser(id string) error
}
