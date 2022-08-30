package usecase

import (
	"crypto/md5"
	"fmt"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"github.com/mateus-sousa-dev/meet-people/app/internal"
	"github.com/mateus-sousa-dev/meet-people/app/tests/mocks"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
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
		Password:        "AbCd123456",
		ConfirmPassword: "AbCd123456",
		Birthday:        839066400,
		Gender:          "M",
	}
	repo := mocks.NewUserRepositoryMock()
	mailRepo := mocks.NewMailRepositoryMock()
	useCase := NewUserUseCase(repo, mailRepo)
	user, err := useCase.CreateUser(userDto)
	assert.Nil(t, err)
	assert.Equal(t, "Mateus", user.FirstName)
	assert.Equal(t, "Silva", user.LastName)
	assert.Equal(t, "mateus@gmail.com", user.Email)
	assert.Equal(t, "", user.MobileNumber)
	assert.Equal(t, 839066400, user.Birthday)
	assert.Equal(t, "M", user.Gender)
	assert.Equal(t, 0, user.Active)
	timeNow := internal.Now().UTC().Unix()
	assert.Equal(t, &timeNow, user.CreatedAt)
	assert.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(time.Now().Unix(), 10)+"mateus@gmail.com"))), user.PathAccountActivation)
	assert.Equal(t, "AbCd123456", user.ConfirmPassword)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("AbCd123456"))
	assert.Nil(t, err)
}

func TestCreateUserInvalid(t *testing.T) {
	userDto := domain.UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "AbCd123456",
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
		Password:        "AbCd123456",
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

func TestValidatePasswordStrengthShort(t *testing.T) {
	password := "Ab14725"
	repo := mocks.NewUserRepositoryMock()
	mailRepo := mocks.NewMailRepositoryMock()
	useCase := NewUserUseCase(repo, mailRepo)
	err := useCase.validatePasswordStrength(password)
	assert.Equal(t, "password is not strong enough", err.Error())
}

func TestValidatePasswordStrengthNotUpper(t *testing.T) {
	password := "abcd147258369"
	repo := mocks.NewUserRepositoryMock()
	mailRepo := mocks.NewMailRepositoryMock()
	useCase := NewUserUseCase(repo, mailRepo)
	err := useCase.validatePasswordStrength(password)
	assert.Equal(t, "password is not strong enough", err.Error())
}

func TestValidatePasswordStrengthNotLower(t *testing.T) {
	password := "ABCD147258369"
	repo := mocks.NewUserRepositoryMock()
	mailRepo := mocks.NewMailRepositoryMock()
	useCase := NewUserUseCase(repo, mailRepo)
	err := useCase.validatePasswordStrength(password)
	assert.Equal(t, "password is not strong enough", err.Error())
}

func TestValidatePasswordStrengthValid(t *testing.T) {
	password := "AbCd147258369"
	repo := mocks.NewUserRepositoryMock()
	mailRepo := mocks.NewMailRepositoryMock()
	useCase := NewUserUseCase(repo, mailRepo)
	err := useCase.validatePasswordStrength(password)
	assert.Nil(t, err)
}
