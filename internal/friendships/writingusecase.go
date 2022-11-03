package friendships

import (
	"errors"
	"github.com/mateus-sousa-dev/meet-people/internal/users"
)

type WritingUseCase interface {
	RequestFriendship(friendshipDto FriendshipDto, loggedUserID int64) (*Friendship, error)
}

type writingUseCase struct {
	repository     Repository
	userRepository users.Repository
}

func NewWritingUseCase(repository Repository, userRepository users.Repository) WritingUseCase {
	return &writingUseCase{repository: repository, userRepository: userRepository}
}

func (w *writingUseCase) RequestFriendship(friendshipDto FriendshipDto, loggedUserID int64) (*Friendship, error) {
	if loggedUserID != friendshipDto.RequesterID {
		return nil, errors.New("you should be requester of friendship you request")
	}
	friendship, err := NewFriendshipRequest(friendshipDto)
	if err != nil {
		return nil, err
	}
	requestedUser := w.userRepository.FindUserById(friendship.RequestedID)
	if requestedUser == nil {
		return nil, errors.New("user requested not found")
	}
	storedFriendship := w.repository.GetFriendship(friendship)
	if storedFriendship != nil {
		return nil, errors.New("friendship was already requested")
	}
	friendship, err = w.repository.RequestFriendship(friendship)
	if err != nil {
		return nil, err
	}
	return friendship, nil
}
