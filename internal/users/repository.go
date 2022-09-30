package users

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) (*User, error)
	FindUserByEmail(email string) *User
	FindUserByPathAccountActivation(path string) *User
	ActivateAccount(user *User) *User
	ChangePassword(password string, userID int64) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (u *repository) CreateUser(user *User) (*User, error) {
	tx := u.db.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *repository) FindUserByEmail(email string) *User {
	var user User
	u.db.Where(&domain.User{Email: email}).First(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}

func (u *repository) FindUserByPathAccountActivation(path string) *User {
	var user User
	u.db.Where(&domain.User{PathAccountActivation: path}).First(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}

func (u *repository) ActivateAccount(user *User) *User {
	u.db.Model(&user).UpdateColumns(user)
	return user
}

func (u *repository) ChangePassword(password string, userID int64) error {
	tx := u.db.Model(&User{}).Where("id = ?", userID).Update("password", password)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
