package usecase

import (
	"testing"
	"time"

	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"github.com/mateus-sousa-dev/meet-people/app/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	userDto := domain.UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}
	expectedUser := &domain.User{
		ID:              1,
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        time.Unix(839066400, 0),
		Gender:          "M",
	}
	repo := mocks.NewUserRepositoryMock()
	useCase := NewUserUseCase(repo)
	user, err := useCase.CreateUser(userDto)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestCreateUserInvalid(t *testing.T) {
	userDto := domain.UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "invalid",
		Birthday:        839066400,
		Gender:          "M",
	}
	repo := mocks.NewUserRepositoryMock()
	useCase := NewUserUseCase(repo)
	user, err := useCase.CreateUser(userDto)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}
