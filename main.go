package main

import (
	"github.com/labstack/echo/v4"
	"github.com/martinsaporiti/golang_mocks_sample/repositories"
	"github.com/martinsaporiti/golang_mocks_sample/server"
	"github.com/martinsaporiti/golang_mocks_sample/services"
)

func main() {

	repo := repositories.NewUser()
	service := services.NewUser(repo)
	server := server.NewEchoServer(service)
	e := echo.New()

	e.POST("/users", server.SaveUserHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
