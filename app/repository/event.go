package repository

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"github.com/streadway/amqp"
)

type EventRepository struct {
	channel *amqp.Channel
}

func NewEventRepository(channel *amqp.Channel) *EventRepository {
	return &EventRepository{channel: channel}
}

func (e *EventRepository) PublishEvent(event *domain.Event) error {
	err := e.channel.Publish(
		event.ExchangeName,
		event.ExchangeKey,
		false,
		false,
		amqp.Publishing{
			ContentType: event.ContentType,
			Body:        []byte(event.Body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
