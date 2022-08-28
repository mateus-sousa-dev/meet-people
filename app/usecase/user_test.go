package usecase

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"github.com/mateus-sousa-dev/meet-people/app/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
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
		Birthday:        839066400,
		Gender:          "M",
		Active:          0,
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
	assert.Equal(t, "passwords are different", err.Error())
	assert.Nil(t, user)
}

func TestCreateUserEmailAlreadyExists(t *testing.T) {
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
	repo := mocks.NewUserRepoEmailAlreadyExistsMock()
	useCase := NewUserUseCase(repo)
	user, err := useCase.CreateUser(userDto)
	assert.Equal(t, "email already exists", err.Error())
	assert.Nil(t, user)
}
