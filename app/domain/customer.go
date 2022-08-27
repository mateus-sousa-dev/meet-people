package domain

import (
	"errors"
	"time"
)

type Customer struct {
	ID              string    `valid:"uuid"`
	FirstName       string    `valid:"notnull"`
	LastName        string    `valid:"notnull"`
	Email           string    `valid:"notnull"`
	MobileNumber    string    `valid:"-"`
	Password        string    `valid:"notnull"`
	ConfirmPassword string    `valid:"notnull"`
	Birthday        time.Time `valid:"notnull"`
	Gender          string    `valid:"notnull"`
}

type CustomerDto struct {
	FirstName       string
	LastName        string
	Email           string
	MobileNumber    string
	Password        string
	ConfirmPassword string
	Birthday        int64
	Gender          string
}

func NewCustomer(CustomerDto CustomerDto) (*Customer, error) {
	customer := &Customer{
		FirstName:       CustomerDto.FirstName,
		LastName:        CustomerDto.LastName,
		Email:           CustomerDto.Email,
		MobileNumber:    CustomerDto.MobileNumber,
		Password:        CustomerDto.Password,
		ConfirmPassword: CustomerDto.ConfirmPassword,
		Birthday:        time.Unix(CustomerDto.Birthday, 0),
		Gender:          CustomerDto.Gender,
	}

	err := customer.Validate()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) Validate() error {
	if c.FirstName == "" {
		return errors.New("first name is required")
	}
	if c.LastName == "" {
		return errors.New("last name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	if c.Password == "" {
		return errors.New("password is required")
	}
	if c.ConfirmPassword != c.Password {
		return errors.New("passwords are different")
	}
	if c.Birthday.Unix() == 0 {
		return errors.New("birthday is required")
	}
	if c.Gender == "" {
		return errors.New("gender is required")
	}
	return nil
}
