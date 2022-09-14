package mocks

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
)

type UserRepositoryMock struct{}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (u *UserRepositoryMock) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = 1
	return user, nil
}

func (u *UserRepositoryMock) FindUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (u *UserRepositoryMock) FindUserByPathAccountActivation(path string) *domain.User {
	return nil
}

func (u *UserRepositoryMock) ActivateAccount(user *domain.User) *domain.User {
	return nil
}

type UserRepoEmailAlreadyExistsMock struct{}

func NewUserRepoEmailAlreadyExistsMock() *UserRepoEmailAlreadyExistsMock {
	return &UserRepoEmailAlreadyExistsMock{}
}

func (u *UserRepoEmailAlreadyExistsMock) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = 1
	return user, nil
}

func (u *UserRepoEmailAlreadyExistsMock) FindUserByEmail(email string) (*domain.User, error) {
	return &domain.User{
		ID:              1,
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}, nil
}

func (u *UserRepoEmailAlreadyExistsMock) FindUserByPathAccountActivation(path string) *domain.User {
	return nil
}

func (u *UserRepoEmailAlreadyExistsMock) ActivateAccount(user *domain.User) *domain.User {
	return nil
}
