package domain

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type UserUseCase interface {
	CreateUser(dto UserDto) (*User, error)
	ActivateAccount(path string) error
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	FindUserByEmail(email string) *User
	FindUserByPathAccountActivation(path string) *User
	ActivateAccount(user *User) *User
}

type User struct {
	ID                    int64
	FirstName             string
	LastName              string
	Email                 string
	MobileNumber          string
	Password              string
	ConfirmPassword       string `gorm:"-"`
	Birthday              int
	Gender                string
	Active                int
	PathAccountActivation string
}

type UserDto struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	MobileNumber    string `json:"mobile_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Birthday        int    `json:"birthday"`
	Gender          string `json:"gender"`
}

func NewUser(userDto UserDto) (*User, error) {
	user := &User{
		FirstName:       userDto.FirstName,
		LastName:        userDto.LastName,
		Email:           userDto.Email,
		MobileNumber:    userDto.MobileNumber,
		Password:        userDto.Password,
		ConfirmPassword: userDto.ConfirmPassword,
		Birthday:        userDto.Birthday,
		Gender:          userDto.Gender,
		Active:          0,
	}
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	user.GeneratePathAccountActivation()
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
	if u.Birthday == 0 {
		return errors.New("birthday is required")
	}
	if u.Gender == "" {
		return errors.New("gender is required")
	}
	return nil
}

func (u *User) GeneratePathAccountActivation() {
	u.PathAccountActivation = fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(time.Now().Unix(), 10)+u.Email)))
}

func (u *User) Activate() error {
	if u.Active == 1 {
		return errors.New("account already is active")
	}
	u.Active = 1
	return nil
}
