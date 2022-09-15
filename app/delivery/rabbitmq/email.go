package rabbitmq

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"github.com/streadway/amqp"
	"os"
)

type EmailDelivery struct {
	rabbitmqChannel *amqp.Channel
	emailUsecase    domain.EmailUseCase
}

func NewEmailDelivery(ch *amqp.Channel, emailUseCase domain.EmailUseCase) *EmailDelivery {
	return &EmailDelivery{
		rabbitmqChannel: ch,
		emailUsecase:    emailUseCase,
	}
}

func (d *EmailDelivery) StartConsume() error {
	msgs, err := d.rabbitmqChannel.Consume(
		os.Getenv("RABBITMQ_QUEUE_NAME"),
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for {
		msg, ok := <-msgs
		if !ok {
			return nil
		}
		err := d.emailUsecase.SendEmail(msg.Body)
		if err != nil {
			err = msg.Nack(true, true)
		}
		err = msg.Ack(true)
		if err != nil {
			return err
		}
	}
}

func (d *EmailDelivery) StartConsume2() error {
	msgs, err := d.rabbitmqChannel.Consume(
		"email2",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for {
		msg, ok := <-msgs
		if !ok {
			return nil
		}
		err := d.emailUsecase.SendEmail(msg.Body)
		if err != nil {
			err = msg.Nack(true, true)
		}
		err = msg.Ack(true)
		if err != nil {
			return err
		}
	}
}
