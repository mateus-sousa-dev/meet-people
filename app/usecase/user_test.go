package usecase

import (
	"crypto/md5"
	"fmt"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"github.com/mateus-sousa-dev/meet-people/app/tests/mocks"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
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
		ID:                    1,
		FirstName:             "Mateus",
		LastName:              "Silva",
		Email:                 "mateus@gmail.com",
		MobileNumber:          "",
		Password:              "123456",
		ConfirmPassword:       "123456",
		Birthday:              839066400,
		Gender:                "M",
		Active:                0,
		PathAccountActivation: fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(time.Now().Unix(), 10)+"mateus@gmail.com"))),
	}
	repo := mocks.NewUserRepositoryMock()
	mailRepo := mocks.NewMailRepositoryMock()
	useCase := NewUserUseCase(repo, mailRepo)
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
	mailRepo := mocks.NewMailRepositoryMock()
	useCase := NewUserUseCase(repo, mailRepo)
	user, err := useCase.CreateUser(userDto)
	assert.Equal(t, "ConfirmPassword: Should be equal to the Password field", err.Error())
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
	mailRepo := mocks.NewMailRepositoryMock()
	useCase := NewUserUseCase(repo, mailRepo)
	user, err := useCase.CreateUser(userDto)
	assert.Equal(t, "email already exists", err.Error())
	assert.Nil(t, user)
}
