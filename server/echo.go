package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func (es *EchoServer) SaveUserHandler(c echo.Context) error {
	type requestType struct {
		Name string `json:"name"`
	}

	var request requestType

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request to create user",
		})
	}

	if request.Name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request to create user, user name is required... :(",
		})
	}

	id, err := es.service.Save("name")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
}
