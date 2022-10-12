package delivery

import (
	"github.com/kapralovs/simple-test-api/internal/users"
	"github.com/labstack/echo/v4"
)

func ResgisterHTTPEndpoints(router *echo.Echo, uc users.Usecase) {
	h := NewHandler(uc)
	router.GET("/get", h.GetUsers)
	router.POST("/edit/:id", h.EditUser)
	router.DELETE("/delete/:id", h.DeleteUser)
}
