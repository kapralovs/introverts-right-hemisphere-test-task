package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kapralovs/introverts-right-hemisphere-test-task/internal/users"
)

func ResgisterHTTPEndpoints(router *fiber.App, uc users.Usecase) {
	handler := New(uc)
	router.Get("/get", handler.Get)
}
