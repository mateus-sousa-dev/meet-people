package domain

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

type Email interface {
	SendAccountActivationEmail(emailTo string) error
}

type EmailSender struct {
}

func (e *EmailSender) SendAccountActivationEmail(emailTo string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "no-reply@meetpeople.com")
	msg.SetHeader("To", emailTo)
	msg.SetHeader("Subject", "Link de ativação da conta")
	msg.SetBody("text/plain", "Clique no link a seguir para ativar sua conta: ")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}
	n := gomail.NewDialer(os.Getenv("SMTP_HOST"), smtpPort, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))
	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
