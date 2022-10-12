package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kapralovs/simple-test-api/internal/server"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err)
	}

	router := echo.New()
	server := server.New(router)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
