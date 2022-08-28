package domain

type MailRepository interface {
	SendMail(emailSender *EmailSender) error
}

type EmailSender struct {
	From        string
	To          string
	Subject     string
	ContentType string
	Body        string
}

type MailSenderDto struct {
	From        string
	To          string
	Subject     string
	ContentType string
	Body        string
}

func NewEmailSender(mailDto MailSenderDto) *EmailSender {
	return &EmailSender{
		From:        mailDto.From,
		To:          mailDto.To,
		Subject:     mailDto.Subject,
		ContentType: mailDto.ContentType,
		Body:        mailDto.Body,
	}
}
