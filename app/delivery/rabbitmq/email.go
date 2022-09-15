package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"
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
		var eventEmailDto domain.EventEmailDto
		err := json.Unmarshal(msg.Body, &eventEmailDto)
		if err != nil {
			err = msg.Nack(true, true)
		}
		err = d.handleUseCase(eventEmailDto)
		if err != nil {
			err = msg.Nack(true, true)
		}
		err = msg.Ack(true)
		if err != nil {
			return err
		}
	}
}

func (d *EmailDelivery) handleUseCase(eventEmailDto domain.EventEmailDto) error {
	if eventEmailDto.Type == "activate-account" {
		return d.emailUsecase.SendAccountActivationEmail(eventEmailDto)
	} else if eventEmailDto.Type == "reset-password" {
		return d.emailUsecase.SendPasswordResetEmail(eventEmailDto)
	}
	fmt.Println("invalido")
	return errors.New("invalid event structure")
}
