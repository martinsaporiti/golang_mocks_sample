package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/martinsaporiti/golang_mocks_sample/models"
	"github.com/martinsaporiti/golang_mocks_sample/services"
)

type EchoServer struct {
	service services.User
}

func NewEchoServer(service services.User) *EchoServer {
	return &EchoServer{
		service: service,
	}
}

func (es *EchoServer) SaveUserHandler(ctx echo.Context) error {

	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request to create user",
		})
	}

	if user.Name == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request to create user, user name is required... :(",
		})
	}

	id, err := es.service.Save("name")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
}
