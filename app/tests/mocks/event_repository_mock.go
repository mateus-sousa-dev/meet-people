package mocks

import "github.com/mateus-sousa-dev/meet-people/app/domain"

type EventRepositoryMock struct{}

func NewEventRepositoryMock() *EventRepositoryMock {
	return &EventRepositoryMock{}
}
func (m *EventRepositoryMock) PublishEvent(body *domain.Event) error {
	return nil
}
