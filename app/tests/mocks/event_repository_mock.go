package mocks

type EventRepositoryMock struct{}

func NewEventRepositoryMock() *EventRepositoryMock {
	return &EventRepositoryMock{}
}
func (m *EventRepositoryMock) PublishEvent(body string) error {
	return nil
}
