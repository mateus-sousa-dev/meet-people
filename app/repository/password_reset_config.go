package repository

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"gorm.io/gorm"
)

type PasswordResetConfigRepository struct {
	db *gorm.DB
}

func NewPasswordResetConfigRepository(db *gorm.DB) *PasswordResetConfigRepository {
	return &PasswordResetConfigRepository{db: db}
}

func (p *PasswordResetConfigRepository) FindPasswordResetConfigByUser(userID int64) *domain.PasswordResetConfig {
	var passwordResetConfig domain.PasswordResetConfig
	p.db.Where(&domain.PasswordResetConfig{UsersID: userID}).First(&passwordResetConfig)
	if passwordResetConfig.ID == 0 {
		return nil
	}
	return &passwordResetConfig
}

func (p *PasswordResetConfigRepository) CreatePasswordResetConfig(passwordResetConfig *domain.PasswordResetConfig) (*domain.PasswordResetConfig, error) {
	tx := p.db.Create(passwordResetConfig)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return passwordResetConfig, nil
}

func (p *PasswordResetConfigRepository) UpdatePasswordResetConfig(passwordResetConfig *domain.PasswordResetConfig) (*domain.PasswordResetConfig, error) {
	tx := p.db.Model(&passwordResetConfig).UpdateColumns(passwordResetConfig)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return passwordResetConfig, nil
}

func (p *PasswordResetConfigRepository) FindPasswordResetConfigByUrl(url string) *domain.PasswordResetConfig {
	var passwordResetConfig domain.PasswordResetConfig
	p.db.Where(&domain.PasswordResetConfig{Url: url}).First(&passwordResetConfig)
	if passwordResetConfig.ID == 0 {
		return nil
	}
	return &passwordResetConfig
}
