package repository

import "github.com/streadway/amqp"

type EventRepository struct {
	channel *amqp.Channel
}

func NewEventRepository(channel *amqp.Channel) *EventRepository {
	return &EventRepository{channel: channel}
}

func (e *EventRepository) PublishEvent(body string) error {
	err := e.channel.Publish("meet-people", "email", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	if err != nil {
		return err
	}
	return nil
}
