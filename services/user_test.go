package services

import (
	"errors"
	"testing"

	mocks "github.com/martinsaporiti/golang_mocks_sample/mocks/repositories"
	"github.com/martinsaporiti/golang_mocks_sample/models"
	"github.com/stretchr/testify/assert"
)

func Test_SaveUser(t *testing.T) {
	// repo := mocks.NewUser(t)
	// service := NewUser(repo)
	// user := models.User{
	// 	Name: "sapo",
	// }
	// repo.On("Save", user).Return(1, nil)

	// id, err := service.Save("sapo")
	// assert.NoError(t, err)
	// assert.Equal(t, 1, id)
	// repo.AssertNumberOfCalls(t, "Save", 1)
	// repo.AssertExpectations(t)

	type wanted struct {
		id  int
		err error
	}

	tests := map[string]struct {
		input string
		repo  func() *mocks.User
		want  wanted
	}{
		"should save user successfuly": {
			input: "sapo",
			repo: func() *mocks.User {
				repo := mocks.NewUser(t)
				user := models.User{
					Name: "sapo",
				}
				repo.On("Save", user).Return(1, nil)
				return repo
			},
			want: wanted{
				id:  1,
				err: nil,
			},
		},

		"should not save the user successfuly": {
			input: "sapo",
			repo: func() *mocks.User {
				repo := mocks.NewUser(t)
				user := models.User{
					Name: "sapo",
				}
				repo.On("Save", user).Return(0, errors.New("Error saving the user"))
				return repo
			},
			want: wanted{
				id:  0,
				err: errors.New("Error saving the user"),
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			repo := tc.repo()
			service := NewUser(repo)
			id, err := service.Save(tc.input)
			if tc.want.err != nil {
				assert.Error(t, err)
			}
			assert.Equal(t, tc.want.id, id)
			repo.AssertExpectations(t)
			repo.AssertNumberOfCalls(t, "Save", 1)
		})
	}
}
