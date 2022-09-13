package domain

type MailRepository interface {
	SendMail(emailSender *MailSender) error
}

type EmailUseCase interface {
	SendEmail(msgBody []byte) error
}

type MailSender struct {
	From        string
	To          string
	Subject     string
	ContentType string
	Body        string
}
