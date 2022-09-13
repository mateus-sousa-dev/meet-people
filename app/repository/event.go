package repository

import (
	"github.com/streadway/amqp"
	"os"
)

type EventRepository struct {
	channel *amqp.Channel
}

func NewEventRepository(channel *amqp.Channel) *EventRepository {
	return &EventRepository{channel: channel}
}

func (e *EventRepository) PublishEvent(body string) error {
	err := e.channel.Publish(
		os.Getenv("RABBITMQ_EXCHANGE_NAME"),
		os.Getenv("RABBITMQ_EXCHANGE_KEY"),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
