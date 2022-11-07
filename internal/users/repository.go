package users

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) (*User, error)
	FindUserById(userID int64) *User
	FindUserByEmail(email string) *User
	FindUserByPathAccountActivation(path string) *User
	ActivateAccount(user *User) *User
	ChangePassword(password string, userID int64) error
	GetMyFriends(loggedUserID int64) []*User
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

func (u *repository) FindUserById(userID int64) *User {
	var user User
	u.db.Where(&User{ID: userID}).First(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}

func (u *repository) FindUserByEmail(email string) *User {
	var user User
	u.db.Where(&User{Email: email}).First(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}

func (u *repository) FindUserByPathAccountActivation(path string) *User {
	var user User
	u.db.Where(&User{PathAccountActivation: path}).First(&user)
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

func (u *repository) GetMyFriends(loggedUserID int64) []*User {
	var users []*User
	u.db.Joins(
		"INNER JOIN DFSGSDGDS ON users.id = friendships.requester_id OR users.id = friendships.requested_id",
	).Find(users)

	return users
}
