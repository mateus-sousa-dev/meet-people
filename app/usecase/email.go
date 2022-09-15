package usecase

import (
	"encoding/json"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"os"
)

type EmailUseCase struct {
	mailRepository domain.MailRepository
}

func NewEmailUseCase(mailRepository domain.MailRepository) *EmailUseCase {
	return &EmailUseCase{mailRepository: mailRepository}
}

func (e *EmailUseCase) SendAccountActivationEmail(eventEmailDto domain.EventEmailDto) error {
	eventBodyJson, err := json.Marshal(eventEmailDto.Body)
	if err != nil {
		return err
	}
	var body *domain.ActivateAccountBody
	err = json.Unmarshal(eventBodyJson, &body)
	if err != nil {
		return err
	}
	urlAccountActivation := os.Getenv("APP_URL") + "/activate-account" + "/" + body.UrlAccountActivation
	mailSender := &domain.MailSender{
		From:        "no-reply@meetpeople.com",
		To:          body.Email,
		Subject:     "Link de ativação",
		ContentType: "text/html",
		Body:        "Clique no link para ativar a sua conta: <a href=\"" + urlAccountActivation + "\">" + urlAccountActivation + "</a>",
	}
	err = e.mailRepository.SendMail(mailSender)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailUseCase) SendPasswordResetEmail(eventEmailDto domain.EventEmailDto) error {
	eventBodyJson, err := json.Marshal(eventEmailDto.Body)
	if err != nil {
		return err
	}
	var body *domain.ResetPasswordBody
	err = json.Unmarshal(eventBodyJson, &body)
	if err != nil {
		return err
	}
	urlPasswordReset := os.Getenv("APP_URL") + "/reset-password" + "/" + body.UrlPasswordReset
	mailSender := &domain.MailSender{
		From:        "no-reply@meetpeople.com",
		To:          body.Email,
		Subject:     "Link para nova senha",
		ContentType: "text/html",
		Body:        "Clique no link para resetar sua senha: <a href=\"" + urlPasswordReset + "\">" + urlPasswordReset + "</a>",
	}
	err = e.mailRepository.SendMail(mailSender)
	if err != nil {
		return err
	}
	return nil
}
