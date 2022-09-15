package domain

import "os"

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

func NewAccountActivationMailSender(body *ActivateAccountBody) *MailSender {
	urlAccountActivation := os.Getenv("APP_URL") + "/activate-account" + "/" + body.UrlAccountActivation
	return &MailSender{
		From:        "no-reply@meetpeople.com",
		To:          body.Email,
		Subject:     "Link de ativação",
		ContentType: "text/html",
		Body:        "Clique no link para ativar a sua conta: <a href=\"" + urlAccountActivation + "\">" + urlAccountActivation + "</a>",
	}
}

func NewPasswordResetMailSender(body *ResetPasswordBody) *MailSender {
	urlPasswordReset := os.Getenv("APP_URL") + "/reset-password" + "/" + body.UrlPasswordReset
	return &MailSender{
		From:        "no-reply@meetpeople.com",
		To:          body.Email,
		Subject:     "Link para nova senha",
		ContentType: "text/html",
		Body:        "Clique no link para resetar sua senha: <a href=\"" + urlPasswordReset + "\">" + urlPasswordReset + "</a>",
	}
}
