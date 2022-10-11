package delivery

import (
	"github.com/kapralovs/introverts-right-hemisphere-test-task/internal/users"
)

type Handler struct {
	uc users.Usecase
}

func New(uc users.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// func (h *Handler) Get(c *fiber.Ctx) error    {}
// func (h *Handler) Delete(c *fiber.Ctx) error {}
// func (h *Handler) Edit(c *fiber.Ctx) error   {}
