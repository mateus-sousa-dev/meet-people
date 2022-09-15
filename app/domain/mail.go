package domain

type MailRepository interface {
	SendMail(emailSender *MailSender) error
}

type EmailUseCase interface {
	SendAccountActivationEmail(EventEmailDto) error
	SendPasswordResetEmail(EventEmailDto) error
}

type EventEmailDto struct {
	Type string
	Body map[string]interface{}
}

type ActivateAccountBody struct {
	Email                string `json:"email"`
	UrlAccountActivation string `json:"urlAccountActivation"`
}

type ResetPasswordBody struct {
	Email            string `json:"email"`
	UrlPasswordReset string `json:"urlPasswordReset"`
}

type MailSender struct {
	From        string
	To          string
	Subject     string
	ContentType string
	Body        string
}
