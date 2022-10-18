package services

import (
	"github.com/martinsaporiti/golang_mocks_sample/models"
	"github.com/martinsaporiti/golang_mocks_sample/repositories"
)

type user struct {
	repo repositories.User
}

func NewUser(repo repositories.User) *user {
	return &user{
		repo: repo,
	}
}

func (us *user) Save(name string) (int, error) {
	return us.repo.Save(models.User{
		Name: name,
	})
}
