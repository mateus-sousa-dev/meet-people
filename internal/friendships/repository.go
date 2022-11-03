package friendships

import "gorm.io/gorm"

type Repository interface {
	GetFriendship(friendship *Friendship) *Friendship
	RequestFriendship(friendship *Friendship) (*Friendship, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetFriendship(friendship *Friendship) *Friendship {
	var storedFriendship Friendship
	r.db.Where(&Friendship{RequesterID: friendship.RequesterID, RequestedID: friendship.RequestedID}).First(&storedFriendship)
	if storedFriendship.ID == 0 {
		return nil
	}
	return &storedFriendship
}

func (r *repository) RequestFriendship(friendship *Friendship) (*Friendship, error) {
	tx := r.db.Create(friendship)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return friendship, nil
}
