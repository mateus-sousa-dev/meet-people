package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
		Active:          0,
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
	assert.Equal(t, "first name is required", err.Error())
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
	assert.Equal(t, "last name is required", err.Error())
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
	assert.Equal(t, "email is required", err.Error())
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
	assert.Equal(t, "password is required", err.Error())
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
	assert.Equal(t, "passwords are different", err.Error())
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
	assert.Equal(t, "birthday is required", err.Error())
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
	assert.Equal(t, "gender is required", err.Error())
}
