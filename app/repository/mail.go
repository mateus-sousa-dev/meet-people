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
func (m *MailRepository) SendMail(mailSender *domain.MailSender) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mailSender.From)
	msg.SetHeader("To", mailSender.To)
	msg.SetHeader("Subject", mailSender.Subject)
	msg.SetBody(mailSender.ContentType, mailSender.Body)
	if err := m.smtpDialer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
