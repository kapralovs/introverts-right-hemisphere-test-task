package server

import "github.com/gofiber/fiber/v2"

type server struct {
	router *fiber.App
}

func New(a *fiber.App) *server {
	return &server{
		router: a,
	}
}

func (s *server) Run() error {
	// mongoClient:=
	// mongoRepo:=
	return s.router.Listen(":8080")
}
