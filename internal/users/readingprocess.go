package users

type ReadingUseCase interface {
	GetMyFriends(loggedUserID int64) []*User
}

type readingUseCase struct {
	repo Repository
}

func NewReadingUseCase(repo Repository) ReadingUseCase {
	return &readingUseCase{repo: repo}
}

func (r *readingUseCase) GetMyFriends(loggedUserID int64) []*User {
	return r.repo.GetMyFriends(loggedUserID)
}
