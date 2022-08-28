package mocks

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
)

type MailRepositoryMock struct{}

func NewMailRepositoryMock() *MailRepositoryMock {
	return &MailRepositoryMock{}
}
func (m *MailRepositoryMock) SendMail(emailSender *domain.MailSender) error {
	return nil
}
