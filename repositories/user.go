package repositories

import "github.com/martinsaporiti/golang_mocks_sample/models"

type User interface {
	Save(user models.User) (int, error)
}
