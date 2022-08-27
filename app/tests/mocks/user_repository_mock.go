package mocks

import "github.com/mateus-sousa-dev/meet-people/app/domain"

type UserRepositoryMock struct{}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (u *UserRepositoryMock) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = 1
	return user, nil
}
