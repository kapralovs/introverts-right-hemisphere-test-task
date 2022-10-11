package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/introverts-right-hemisphere-test-task/internal/server"
)

func main() {
	router := fiber.New()
	server := server.New(router)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
