package usecase

import (
	"encoding/json"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
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
	mailSender := domain.NewAccountActivationMailSender(body)
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
	mailSender := domain.NewPasswordResetMailSender(body)
	err = e.mailRepository.SendMail(mailSender)
	if err != nil {
		return err
	}
	return nil
}
