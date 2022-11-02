package events

import (
	"fmt"
	"os"
)

type Event struct {
	Body         string
	ContentType  string
	ExchangeKey  string
	ExchangeName string
}

func NewActivateAccountEvent(email, urlAccountActivation string) *Event {
	return &Event{
		Body:         fmt.Sprintf("{\"type\":\"activate-account\", \"body\" : {\"email\":\"%s\", \"urlAccountActivation\":\"%s\"}}", email, urlAccountActivation),
		ContentType:  "text/plain",
		ExchangeKey:  os.Getenv("RABBITMQ_EXCHANGE_KEY"),
		ExchangeName: os.Getenv("RABBITMQ_EXCHANGE_NAME"),
	}
}

func NewResetPasswordEvent(email, urlPasswordReset string) *Event {
	return &Event{
		Body:         fmt.Sprintf("{\"type\":\"reset-password\", \"body\" : {\"email\":\"%s\", \"urlPasswordReset\":\"%s\"}}", email, urlPasswordReset),
		ContentType:  "text/plain",
		ExchangeKey:  os.Getenv("RABBITMQ_EXCHANGE_KEY"),
		ExchangeName: os.Getenv("RABBITMQ_EXCHANGE_NAME"),
	}
}
