package domain

import (
	"errors"
	"time"
)

type UserUseCase interface {
	CreateUser(dto UserDto) (*User, error)
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
}

type User struct {
	ID              int64
	FirstName       string
	LastName        string
	Email           string
	MobileNumber    string
	Password        string
	ConfirmPassword string
	Birthday        time.Time
	Gender          string
}

type UserDto struct {
	FirstName       string
	LastName        string
	Email           string
	MobileNumber    string
	Password        string
	ConfirmPassword string
	Birthday        int64
	Gender          string
}

func NewUser(userDto UserDto) (*User, error) {
	user := &User{
		FirstName:       userDto.FirstName,
		LastName:        userDto.LastName,
		Email:           userDto.Email,
		MobileNumber:    userDto.MobileNumber,
		Password:        userDto.Password,
		ConfirmPassword: userDto.ConfirmPassword,
		Birthday:        time.Unix(userDto.Birthday, 0),
		Gender:          userDto.Gender,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if u.ConfirmPassword != u.Password {
		return errors.New("passwords are different")
	}
	if u.Birthday.Unix() == 0 {
		return errors.New("birthday is required")
	}
	if u.Gender == "" {
		return errors.New("gender is required")
	}
	return nil
}
