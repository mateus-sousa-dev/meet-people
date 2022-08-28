package domain

type MailRepository interface {
	SendMail(emailSender *MailSender) error
}

type MailSender struct {
	From        string
	To          string
	Subject     string
	ContentType string
	Body        string
}
