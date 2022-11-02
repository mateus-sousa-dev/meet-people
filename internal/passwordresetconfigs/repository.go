package passwordresetconfigs

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindPasswordResetConfigByUser(userID int64) *PasswordResetConfig
	CreatePasswordResetConfig(passwordResetConfig *PasswordResetConfig) (*PasswordResetConfig, error)
	UpdatePasswordResetConfig(passwordResetConfig *PasswordResetConfig) (*PasswordResetConfig, error)
	FindPasswordResetConfigByUrl(url string) *PasswordResetConfig
	ExpireByUse(passwordResetConfigID int64) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (p *repository) FindPasswordResetConfigByUser(userID int64) *PasswordResetConfig {
	var passwordResetConfig PasswordResetConfig
	p.db.Where(&PasswordResetConfig{UsersID: userID}).First(&passwordResetConfig)
	if passwordResetConfig.ID == 0 {
		return nil
	}
	return &passwordResetConfig
}

func (p *repository) CreatePasswordResetConfig(passwordResetConfig *PasswordResetConfig) (*PasswordResetConfig, error) {
	tx := p.db.Create(passwordResetConfig)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return passwordResetConfig, nil
}

func (p *repository) UpdatePasswordResetConfig(passwordResetConfig *PasswordResetConfig) (*PasswordResetConfig, error) {
	tx := p.db.Model(&passwordResetConfig).UpdateColumns(passwordResetConfig)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return passwordResetConfig, nil
}

func (p *repository) FindPasswordResetConfigByUrl(url string) *PasswordResetConfig {
	var passwordResetConfig PasswordResetConfig
	p.db.Where(&PasswordResetConfig{Url: url}).First(&passwordResetConfig)
	if passwordResetConfig.ID == 0 {
		return nil
	}
	return &passwordResetConfig
}

func (p *repository) ExpireByUse(passwordResetConfigID int64) error {
	tx := p.db.Model(&PasswordResetConfig{}).Where("id = ?", passwordResetConfigID).Update("expiration_by_use", 1)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
