package infra

import (
	"github.com/streadway/amqp"
	"os"
	"strconv"
)

func StartRabbitMQ() (*amqp.Channel, error) {
	port, err := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	uri := amqp.URI{
		Scheme:   "amqp",
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     port,
		Username: os.Getenv("RABBITMQ_USERNAME"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
		Vhost:    os.Getenv("RABBITMQ_VHOST"),
	}
	conn, err := amqp.Dial(uri.String())
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		os.Getenv("RABBITMQ_QUEUE_NAME"),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	q2, err := ch.QueueDeclare(
		"email2",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = ch.ExchangeDeclare(os.Getenv("RABBITMQ_EXCHANGE_NAME"),
		os.Getenv("RABBITMQ_EXCHANGE_KIND"),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = ch.QueueBind(q.Name,
		os.Getenv("RABBITMQ_EXCHANGE_KEY"),
		os.Getenv("RABBITMQ_EXCHANGE_NAME"),
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = ch.QueueBind(q2.Name,
		"email2",
		os.Getenv("RABBITMQ_EXCHANGE_NAME"),
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
