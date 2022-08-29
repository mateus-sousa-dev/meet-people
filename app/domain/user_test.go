package domain

import (
	"crypto/md5"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	expectedUser := &User{
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
	user, err := NewUser(userDto)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
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
	assert.Equal(t, "FirstName: Missing required field", err.Error())
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
	assert.Equal(t, "LastName: Missing required field", err.Error())
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
	assert.Equal(t, "Email: Missing required field", err.Error())
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
	assert.Equal(t, "Gender: Missing required field", err.Error())
}
