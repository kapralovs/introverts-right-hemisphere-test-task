package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kapralovs/simple-test-api/internal/server"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err)
	}

	time.Sleep(5 * time.Second)
	server := server.New(echo.New())
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
