package repository

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	tx := u.db.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	tx := u.db.Where(&domain.User{Email: email}).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if user.ID == 0 {
		return nil, nil
	}
	return &user, nil
}

func (u *UserRepository) FindUserByPathAccountActivation(path string) *domain.User {
	var user domain.User
	u.db.Where(&domain.User{PathAccountActivation: path}).First(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}

func (u *UserRepository) ActivateAccount(user *domain.User) *domain.User {
	u.db.Model(&user).UpdateColumns(user)
	return user
}
