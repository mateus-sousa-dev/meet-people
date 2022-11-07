package friendships

import (
	"errors"
	"github.com/mateus-sousa-dev/meet-people/internal/global"
	"github.com/mateus-sousa-dev/meet-people/internal/users"
)

type Friendship struct {
	ID          int64
	RequesterID int64 `gorm:"column:requester_id;type:int REFERENCES users(id);notnull"`
	Requester   *users.User
	RequestedID int64 `gorm:"column:requested_id;type:int REFERENCES users(id);notnull"`
	Requested   *users.User
	RequestedAt int64
	AcceptedAt  int64
	Accepted    int64
}

type FriendshipDto struct {
	ID          int64 `json:"id"`
	RequesterID int64 `json:"requester_id"`
	RequestedID int64 `json:"requested_id"`
}

const (
	PENDING_REQUEST  = 0
	ACCEPTED_REQUEST = 1
)

func NewFriendshipRequest(friendshipDto FriendshipDto) (*Friendship, error) {
	if friendshipDto.RequesterID == friendshipDto.RequestedID {
		return nil, errors.New("requester user id cannot be equal requested user id")
	}
	return &Friendship{
		RequesterID: friendshipDto.RequesterID,
		RequestedID: friendshipDto.RequestedID,
		RequestedAt: global.Now().UTC().Unix(),
	}, nil
}

func (f *Friendship) beAccepted() error {
	if f.Accepted != PENDING_REQUEST {
		return errors.New("friendship request was already accepted")
	}
	f.Accepted = ACCEPTED_REQUEST
	f.AcceptedAt = global.Now().UTC().Unix()
	return nil
}
