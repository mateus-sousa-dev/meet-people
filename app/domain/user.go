package domain

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
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
	ID                    int64  `json:"id" valid:"-"`
	FirstName             string `json:"first_name" valid:"notnull"`
	LastName              string `json:"last_name" valid:"notnull"`
	Email                 string `json:"email" valid:"notnull"`
	MobileNumber          string `json:"mobile_number" valid:"-"`
	Password              string `json:"-" valid:"notnull"`
	ConfirmPassword       string `json:"-" gorm:"-" valid:"notnull"`
	Birthday              int    `json:"birthday" valid:"-"`
	Gender                string `json:"gender" valid:"notnull"`
	Active                int    `json:"active" valid:"-"`
	PathAccountActivation string `json:"-" valid:"-"`
	CreatedAt             int    `json:"-" valid:"-"`
	UpdatedAt             int    `json:"-" valid:"-"`
	DeletedAt             int    `json:"-" valid:"-"`
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

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
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
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	if u.Birthday == 0 {
		return errors.New("Birthday: Missing required field")
	}
	if u.ConfirmPassword != u.Password {
		return errors.New("ConfirmPassword: Should be equal to the Password field")
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
