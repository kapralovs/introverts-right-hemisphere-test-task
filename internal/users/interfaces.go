package users

import "github.com/kapralovs/introverts-right-hemisphere-test-task/internal/models"

type Usecase interface {
	Get() ([]*models.User, error)
	Edit(id int) error
	Delete(id int) error
}

type Repository interface {
	Get() ([]*models.User, error)
	Edit(id int) error
	Delete(id int) error
}
