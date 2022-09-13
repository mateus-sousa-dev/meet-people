package rabbitmq

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"github.com/streadway/amqp"
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
	msgs, err := d.rabbitmqChannel.Consume("email", "", true, false, false, false, nil)
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
	}
	return nil
}
