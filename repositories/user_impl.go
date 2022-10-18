package repositories

import (
	"errors"

	"github.com/martinsaporiti/golang_mocks_sample/models"
)

type user struct {
	count int
}

func NewUser() *user {
	return &user{
		count: 0,
	}
}

func (ur *user) Save(user models.User) (int, error) {
	ur.count += 1
	if ur.count == 3 {
		return 0, errors.New("Error creating user")
	}
	return ur.count, nil
}
