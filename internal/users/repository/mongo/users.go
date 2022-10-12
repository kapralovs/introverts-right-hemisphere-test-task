package mongo

import (
	"context"

	"github.com/kapralovs/simple-test-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

func (r *UserRepository) GetUsers() ([]*models.User, error) {
	cur, err := r.db.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	users := make([]*User, 0)
	for cur.Next(context.Background()) {
		user := new(User)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return toUsers(users), nil
}

func (r *UserRepository) EditUser(data *models.User) error {
	newData := toModel(data)
	update := bson.D{{"$set", bson.D{{"name", newData.Name}}}}
	_, err := r.db.UpdateByID(context.Background(), newData.ID, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := r.db.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

func toModel(u *models.User) *User {
	uid, _ := primitive.ObjectIDFromHex(u.ID)
	model := &User{
		ID:   uid,
		Name: u.Name,
	}

	return model
}

func toUser(u *User) *models.User {
	return &models.User{
		ID:   u.ID.Hex(),
		Name: u.Name,
	}
}

func toUsers(users []*User) []*models.User {
	out := make([]*models.User, len(users))

	for idx, u := range users {
		out[idx] = toUser(u)
	}

	return out
}
