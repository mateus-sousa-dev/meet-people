package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewCustomerValid(t *testing.T) {
	customerDto := CustomerDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}
	expectedCustomer := &Customer{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        time.Unix(839066400, 0),
		Gender:          "M",
	}
	customer, err := NewCustomer(customerDto)
	assert.Nil(t, err)
	assert.Equal(t, expectedCustomer, customer)
}

func TestNewCustomerWithoutFirstName(t *testing.T) {
	customerDto := CustomerDto{
		FirstName:       "",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewCustomer(customerDto)
	assert.NotNil(t, err)
	assert.Equal(t, "first name is required", err.Error())
}

func TestNewCustomerWithoutLastName(t *testing.T) {
	customerDto := CustomerDto{
		FirstName:       "Mateus",
		LastName:        "",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewCustomer(customerDto)
	assert.NotNil(t, err)
	assert.Equal(t, "last name is required", err.Error())
}

func TestNewCustomerWithoutEmail(t *testing.T) {
	customerDto := CustomerDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewCustomer(customerDto)
	assert.NotNil(t, err)
	assert.Equal(t, "email is required", err.Error())
}

func TestNewCustomerWithoutPassword(t *testing.T) {
	customerDto := CustomerDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewCustomer(customerDto)
	assert.NotNil(t, err)
	assert.Equal(t, "password is required", err.Error())
}

func TestNewCustomerDifferentPasswords(t *testing.T) {
	customerDto := CustomerDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "12345",
		Birthday:        839066400,
		Gender:          "M",
	}

	_, err := NewCustomer(customerDto)
	assert.NotNil(t, err)
	assert.Equal(t, "passwords are different", err.Error())
}

func TestNewCustomerWithoutBirthday(t *testing.T) {
	customerDto := CustomerDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        0,
		Gender:          "M",
	}

	_, err := NewCustomer(customerDto)
	assert.NotNil(t, err)
	assert.Equal(t, "birthday is required", err.Error())
}

func TestNewCustomerWithoutGender(t *testing.T) {
	customerDto := CustomerDto{
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "",
	}

	_, err := NewCustomer(customerDto)
	assert.NotNil(t, err)
	assert.Equal(t, "gender is required", err.Error())
}
