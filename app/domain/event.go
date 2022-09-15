package domain

type EventRepository interface {
	PublishEvent(body string) error
	PublishEvent2(body string) error
}
