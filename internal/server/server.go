package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kapralovs/simple-test-api/internal/users"
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
	uc     users.Usecase
}

func New(r *echo.Echo) *server {
	db := initDB()
	mongoRepo := repo.NewUserRepository(db, "users")
	return &server{
		uc: usecase.New(mongoRepo),
	}
}

func (s *server) Run() error {
	r := echo.New()
	s.router = r
	delivery.ResgisterHTTPEndpoints(s.router, s.uc)
	fmt.Println("Server is starting...")
	return s.router.Start(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
}

func initDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://root:example@%s:%s/", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT")))
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
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	return client.Database("testdb")
}
