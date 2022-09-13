package domain

type EventRepository interface {
	PublishEvent(body string) error
}
