package friendships

import (
	"errors"
	"github.com/mateus-sousa-dev/meet-people/internal/global"
)

type Friendship struct {
	ID          int64
	RequesterID int64
	RequestedID int64
	RequestedAt int64
	AcceptedAt  int64
	Accepted    int64
}

type FriendshipDto struct {
	RequesterID int64 `json:"requester_id"`
	RequestedID int64 `json:"requested_id"`
}

func NewFriendshipRequest(friendshipDto FriendshipDto) (*Friendship, error) {
	if friendshipDto.RequesterID == friendshipDto.RequestedID {
		return nil, errors.New("requester user id cannot be equal requested user id")
	}
	return &Friendship{
		ID:          0,
		RequesterID: friendshipDto.RequesterID,
		RequestedID: friendshipDto.RequestedID,
		RequestedAt: global.Now().UTC().Unix(),
	}, nil
}
