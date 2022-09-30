package emails

import (
	"encoding/json"
)

type SendUseCase interface {
	SendAccountActivationEmail(EventEmailDto) error
	SendPasswordResetEmail(EventEmailDto) error
}

type sendUseCase struct {
	repository Repository
}

func NewSendUseCase(repository Repository) SendUseCase {
	return &sendUseCase{repository: repository}
}

func (e *sendUseCase) SendAccountActivationEmail(eventEmailDto EventEmailDto) error {
	eventBodyJson, err := json.Marshal(eventEmailDto.Body)
	if err != nil {
		return err
	}
	var body *ActivateAccountBody
	err = json.Unmarshal(eventBodyJson, &body)
	if err != nil {
		return err
	}
	mailSender := NewAccountActivationMailSender(body)
	err = e.repository.SendMail(mailSender)
	if err != nil {
		return err
	}
	return nil
}

func (e *sendUseCase) SendPasswordResetEmail(eventEmailDto EventEmailDto) error {
	eventBodyJson, err := json.Marshal(eventEmailDto.Body)
	if err != nil {
		return err
	}
	var body *ResetPasswordBody
	err = json.Unmarshal(eventBodyJson, &body)
	if err != nil {
		return err
	}
	mailSender := NewPasswordResetMailSender(body)
	err = e.repository.SendMail(mailSender)
	if err != nil {
		return err
	}
	return nil
}
