package delivery

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/kapralovs/simple-test-api/internal/models"
	"github.com/kapralovs/simple-test-api/internal/users"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	useCase users.Usecase
}

func NewHandler(uc users.Usecase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.useCase.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	json, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	return c.String(http.StatusOK, string(json))
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errors.New("empty id param"))
	}
	err := h.useCase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(200, "deleted!")
}

func (h *Handler) EditUser(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if body == nil {
		return c.JSON(http.StatusBadRequest, errors.New("nil request body"))
	}
	data := new(models.User)
	err = json.Unmarshal(body, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.useCase.EditUser(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, "edited!")
}
