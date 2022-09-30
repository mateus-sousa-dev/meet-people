package emails

import (
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"os"
)

type EmailDelivery struct {
	rabbitmqChannel *amqp.Channel
	sendUsecase     SendUseCase
}

func NewDelivery(ch *amqp.Channel, sendUsecase SendUseCase) *EmailDelivery {
	return &EmailDelivery{
		rabbitmqChannel: ch,
		sendUsecase:     sendUsecase,
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
		var eventEmailDto EventEmailDto
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

func (d *EmailDelivery) handleUseCase(eventEmailDto EventEmailDto) error {
	if eventEmailDto.Type == "activate-account" {
		return d.sendUsecase.SendAccountActivationEmail(eventEmailDto)
	} else if eventEmailDto.Type == "reset-password" {
		return d.sendUsecase.SendPasswordResetEmail(eventEmailDto)
	}
	return errors.New("invalid event structure")
}
