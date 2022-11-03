package friendships

import "gorm.io/gorm"

type Repository interface {
	GetFriendship(friendship *Friendship) *Friendship
	RequestFriendship(friendship *Friendship) (*Friendship, error)
	AcceptFriendship(friendship *Friendship) (*Friendship, error)
	GetFriendshipRequestById(friendshipID int64) *Friendship
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

func (r *repository) GetFriendshipRequestById(friendshipID int64) *Friendship {
	var friendship Friendship
	r.db.Where(&Friendship{ID: friendshipID}).First(&friendship)
	if friendship.ID == 0 {
		return nil
	}
	return &friendship
}

func (r *repository) AcceptFriendship(friendship *Friendship) (*Friendship, error) {
	tx := r.db.Model(
		&Friendship{},
	).Where(
		"id = ?",
		friendship.ID,
	).Updates(friendship)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return friendship, nil
}
