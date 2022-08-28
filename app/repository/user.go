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
	u.db.Create(user)
	return user, nil
}

func (u *UserRepository) FindUserByEmail(email string) *domain.User {
	var user domain.User
	u.db.Where(&domain.User{Email: email}).First(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}
