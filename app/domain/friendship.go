package domain

type Friendship struct {
	ID              string
	Requester       *User
	Requested       *User
	RequestApproved int
	RequestedAt     int
	ApprovedAt      int
}
