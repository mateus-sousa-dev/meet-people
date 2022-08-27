package usecase

import "github.com/mateus-sousa-dev/meet-people/app/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(userDto domain.UserDto) (*domain.User, error) {
	user, err := domain.NewUser(userDto)
	if err != nil {
		return nil, err
	}
	user, err = u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
