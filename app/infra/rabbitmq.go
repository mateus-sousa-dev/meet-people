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

	q, err := ch.QueueDeclare("email", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare("meet-people", "direct", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(q.Name, "email", "meet-people", false, nil)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
