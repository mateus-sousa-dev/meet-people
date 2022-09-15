package usecase

import (
	"errors"
	"fmt"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"github.com/mateus-sousa-dev/meet-people/app/internal"
	"strings"
)

type UserUseCase struct {
	repo            domain.UserRepository
	eventRepository domain.EventRepository
}

func NewUserUseCase(repo domain.UserRepository, eventRepository domain.EventRepository) *UserUseCase {
	return &UserUseCase{repo: repo, eventRepository: eventRepository}
}

func (u *UserUseCase) CreateUser(userDto domain.UserDto) (*domain.User, error) {
	user := u.repo.FindUserByEmail(userDto.Email)
	if user != nil {
		return nil, errors.New("email already exists")
	}
	err := u.validatePasswordStrength(userDto.Password)
	if err != nil {
		return nil, err
	}
	user, err = domain.NewUser(userDto)
	if err != nil {
		return nil, err
	}
	timeNow := internal.Now().UTC().Unix()
	user.CreatedAt = &timeNow
	user.UpdatedAt = &timeNow
	user, err = u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	bodyEvent := fmt.Sprintf("{ \"email\":\"%s\", \"urlAccountActivation\":\"%s\"}", user.Email, user.PathAccountActivation)
	err = u.eventRepository.PublishEvent(bodyEvent)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) ActivateAccount(path string) error {
	user := u.repo.FindUserByPathAccountActivation(path)
	if user == nil {
		return errors.New("user not found")
	}
	err := user.Activate()
	if err != nil {
		return err
	}
	u.repo.ActivateAccount(user)
	return nil
}

func (u *UserUseCase) validatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password is not strong enough")
	}
	if strings.ToLower(password) == password {
		return errors.New("password is not strong enough")
	}
	if strings.ToUpper(password) == password {
		return errors.New("password is not strong enough")
	}
	return nil
}

func (u *UserUseCase) ForgotPassword(email string) error {
	user := u.repo.FindUserByEmail(email)
	if user == nil {
		return errors.New("your search did not return any results")
	}
	return nil
}
