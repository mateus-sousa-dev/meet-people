package users

import (
	"github.com/mateus-sousa-dev/meet-people/internal/events"
	"github.com/mateus-sousa-dev/meet-people/internal/passwordresetconfigs"
)

type UserRepositoryMock struct{}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (u *UserRepositoryMock) CreateUser(user *User) (*User, error) {
	user.ID = 1
	return user, nil
}

func (u *UserRepositoryMock) FindUserByEmail(email string) *User {
	return nil
}

func (u *UserRepositoryMock) FindUserByPathAccountActivation(path string) *User {
	return nil
}

func (u *UserRepositoryMock) ActivateAccount(user *User) *User {
	return nil
}

func (u *UserRepositoryMock) ChangePassword(password string, userID int64) error {
	return nil
}

type UserRepoEmailAlreadyExistsMock struct{}

func NewUserRepoEmailAlreadyExistsMock() *UserRepoEmailAlreadyExistsMock {
	return &UserRepoEmailAlreadyExistsMock{}
}

func (u *UserRepoEmailAlreadyExistsMock) CreateUser(user *User) (*User, error) {
	user.ID = 1
	return user, nil
}

func (u *UserRepoEmailAlreadyExistsMock) FindUserByEmail(email string) *User {
	return &User{
		ID:              1,
		FirstName:       "Mateus",
		LastName:        "Silva",
		Email:           "mateus@gmail.com",
		MobileNumber:    "",
		Password:        "123456",
		ConfirmPassword: "123456",
		Birthday:        839066400,
		Gender:          "M",
	}
}

func (u *UserRepoEmailAlreadyExistsMock) FindUserByPathAccountActivation(path string) *User {
	return nil
}

func (u *UserRepoEmailAlreadyExistsMock) ActivateAccount(user *User) *User {
	return nil
}

func (u *UserRepoEmailAlreadyExistsMock) ChangePassword(password string, userID int64) error {
	return nil
}

type EventRepositoryMock struct{}

func NewEventRepositoryMock() *EventRepositoryMock {
	return &EventRepositoryMock{}
}
func (m *EventRepositoryMock) PublishEvent(body *events.Event) error {
	return nil
}

type PasswordResetConfigRepositoryMock struct{}

func NewPasswordResetConfigRepositoryMock() *PasswordResetConfigRepositoryMock {
	return &PasswordResetConfigRepositoryMock{}
}

func (p *PasswordResetConfigRepositoryMock) FindPasswordResetConfigByUser(userID int64) *passwordresetconfigs.PasswordResetConfig {
	return nil
}

func (p *PasswordResetConfigRepositoryMock) CreatePasswordResetConfig(passwordResetConfig *passwordresetconfigs.PasswordResetConfig) (*passwordresetconfigs.PasswordResetConfig, error) {
	return nil, nil
}

func (p *PasswordResetConfigRepositoryMock) UpdatePasswordResetConfig(passwordResetConfig *passwordresetconfigs.PasswordResetConfig) (*passwordresetconfigs.PasswordResetConfig, error) {
	return nil, nil
}

func (p *PasswordResetConfigRepositoryMock) FindPasswordResetConfigByUrl(url string) *passwordresetconfigs.PasswordResetConfig {
	return nil
}

func (p *PasswordResetConfigRepositoryMock) ExpireByUse(passwordResetConfigID int64) error {
	return nil
}
