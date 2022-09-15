package mocks

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
)

type PasswordResetConfigRepositoryMock struct{}

func NewPasswordResetConfigRepositoryMock() *PasswordResetConfigRepositoryMock {
	return &PasswordResetConfigRepositoryMock{}
}

func (p *PasswordResetConfigRepositoryMock) FindPasswordResetConfigByUser(userID int64) *domain.PasswordResetConfig {
	return nil
}

func (p *PasswordResetConfigRepositoryMock) CreatePasswordResetConfig(passwordResetConfig *domain.PasswordResetConfig) (*domain.PasswordResetConfig, error) {
	return nil, nil
}

func (p *PasswordResetConfigRepositoryMock) UpdatePasswordResetConfig(passwordResetConfig *domain.PasswordResetConfig) (*domain.PasswordResetConfig, error) {
	return nil, nil
}

func (p *PasswordResetConfigRepositoryMock) FindPasswordResetConfigByUrl(url string) *domain.PasswordResetConfig {
	return nil
}

func (p *PasswordResetConfigRepositoryMock) ExpireByUse(passwordResetConfigID int64) error {
	return nil
}
