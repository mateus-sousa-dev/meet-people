package domain

import (
	"crypto/md5"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"testing"
	"time"
)

func TestNewUserValid(t *testing.T) {
	userDto := UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}
	user, err := NewUser(userDto)
	assert.Nil(t, err)
	assert.Equal(t, "Mateus", user.FirstName)
	assert.Equal(t, "Silva", user.LastName)
	assert.Equal(t, "mateus@gmail.com", user.Email)
	assert.Equal(t, "", user.MobileNumber)
	assert.Equal(t, 839066400, user.Birthday)
	assert.Equal(t, "M", user.Gender)
	assert.Equal(t, 0, user.Active)
	assert.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(time.Now().Unix(), 10)+"mateus@gmail.com"))), user.PathAccountActivation)
	assert.Equal(t, "123456", user.ConfirmPassword)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("123456"))
	assert.Nil(t, err)
}

func TestNewUserWithoutFirstName(t *testing.T) {
	userDto := UserDto{
		FirstName:       "",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewUser(userDto)
	assert.NotNil(t, err)
	assert.Equal(t, "first_name: Missing required field", err.Error())
}

func TestNewUserWithoutLastName(t *testing.T) {
	userDto := UserDto{
		FirstName:       "Mateus",
		LastName:        "",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewUser(userDto)
	assert.NotNil(t, err)
	assert.Equal(t, "last_name: Missing required field", err.Error())
}

func TestNewUserWithoutEmail(t *testing.T) {
	userDto := UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewUser(userDto)
	assert.NotNil(t, err)
	assert.Equal(t, "email: Missing required field", err.Error())
}

func TestNewUserWithoutPassword(t *testing.T) {
	userDto := UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewUser(userDto)
	assert.NotNil(t, err)
	assert.Equal(t, "Password: Missing required field", err.Error())
}

func TestNewUserDifferentPasswords(t *testing.T) {
	userDto := UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "12345",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewUser(userDto)
	assert.NotNil(t, err)
	assert.Equal(t, "ConfirmPassword: Should be equal to the Password field", err.Error())
}

func TestNewUserWithoutBirthday(t *testing.T) {
	userDto := UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        0,
		Gender:          "M",
	}

	_, err := NewUser(userDto)
	assert.NotNil(t, err)
	assert.Equal(t, "Birthday: Missing required field", err.Error())
}

func TestNewUserWithoutGender(t *testing.T) {
	userDto := UserDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "",
	}

	_, err := NewUser(userDto)
	assert.NotNil(t, err)
	assert.Equal(t, "gender: Missing required field", err.Error())
}
