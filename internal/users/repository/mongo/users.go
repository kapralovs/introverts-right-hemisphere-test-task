package mongo

import (
	"github.com/kapralovs/introverts-right-hemisphere-test-task/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

func (r *UserRepository) Get() ([]*models.User, error) {
	users := []*models.User{}
	return users, nil
}

func (r *UserRepository) Edit(id int) error {
	return nil
}

func (r *UserRepository) Delete(id int) error {
	return nil
}
