package events

import "github.com/streadway/amqp"

type Repository interface {
	PublishEvent(body *Event) error
}

type repository struct {
	channel *amqp.Channel
}

func NewRepository(channel *amqp.Channel) Repository {
	return &repository{channel: channel}
}

func (e *repository) PublishEvent(event *Event) error {
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
