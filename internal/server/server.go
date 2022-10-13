package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kapralovs/simple-test-api/internal/models"
	"github.com/kapralovs/simple-test-api/internal/users/delivery"
	repo "github.com/kapralovs/simple-test-api/internal/users/repository/mongo"
	"github.com/kapralovs/simple-test-api/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type server struct {
	router *echo.Echo
}

func New(r *echo.Echo) *server {
	return &server{
		router: echo.New(),
	}
}

func (s *server) Run() error {
	db := initDB()
	mongoRepo := repo.NewUserRepository(db, "users")
	uc := usecase.New(mongoRepo)
	err := fillDB(db)
	if err != nil {
		fmt.Println(err)
	}
	delivery.ResgisterHTTPEndpoints(s.router, uc)
	return s.router.Start(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
}

func initDB() *mongo.Database {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	fmt.Println("HOST: ", host)
	fmt.Println("PORT: ", port)
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://root:example@%s:%s/?connect=direct", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT")))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.Background()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	return client.Database("simple-api")
}

func fillDB(db *mongo.Database) error {
	users := []*models.User{
		{
			ID:   "1",
			Name: "Jeff",
		},
		{
			ID:   "2",
			Name: "Scott",
		},
	}
	for _, u := range users {
		collection := db.Collection("users")
		_, err := collection.InsertOne(context.TODO(), u)
		if err != nil {
			return err
		}
	}
	return nil
}
