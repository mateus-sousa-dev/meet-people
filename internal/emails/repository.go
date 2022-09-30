package emails

import (
	"gopkg.in/gomail.v2"
)

type Repository interface {
	SendMail(emailSender *MailSender) error
}

type repository struct {
	smtpDialer *gomail.Dialer
}

func NewRepository(smtpDialer *gomail.Dialer) Repository {
	return &repository{smtpDialer: smtpDialer}
}

func (m *repository) SendMail(mailSender *MailSender) error {
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
