package repository

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"gopkg.in/gomail.v2"
)

type MailRepository struct {
	smtpDialer *gomail.Dialer
}

func NewMailRepository(smtpDialer *gomail.Dialer) *MailRepository {
	return &MailRepository{smtpDialer: smtpDialer}
}
func (m *MailRepository) SendMail(emailSender *domain.EmailSender) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", emailSender.From)
	msg.SetHeader("To", emailSender.To)
	msg.SetHeader("Subject", emailSender.Subject)
	msg.SetBody(emailSender.ContentType, emailSender.Body)
	if err := m.smtpDialer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
