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

type Body struct {
	Email                string `json:"email"`
	UrlAccountActivation string `json:"urlAccountActivation"`
}

func (e *EmailUseCase) SendEmail(msg []byte) error {
	var body Body
	err := json.Unmarshal(msg, &body)
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

func (e *EmailUseCase) SendEmail2(msg []byte) error {
	var body Body
	err := json.Unmarshal(msg, &body)
	if err != nil {
		return err
	}
	urlAccountActivation := os.Getenv("APP_URL") + "/activate-account" + "/" + body.UrlAccountActivation
	mailSender := &domain.MailSender{
		From:        "no-reply@meetpeople.com",
		To:          body.Email,
		Subject:     "Link de ativação",
		ContentType: "text/html",
		Body:        "Clique no link para resetar sua senha: <a href=\"" + urlAccountActivation + "\">" + urlAccountActivation + "</a>",
	}
	err = e.mailRepository.SendMail(mailSender)
	if err != nil {
		return err
	}
	return nil
}
