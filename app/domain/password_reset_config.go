package domain

import (
	"crypto/md5"
	"fmt"
	"github.com/mateus-sousa-dev/meet-people/app/internal"
	"strconv"
	"time"
)

type PasswordResetConfigRepository interface {
	FindPasswordResetConfigByUser(userID int64) *PasswordResetConfig
	CreatePasswordResetConfig(passwordResetConfig *PasswordResetConfig) (*PasswordResetConfig, error)
	UpdatePasswordResetConfig(passwordResetConfig *PasswordResetConfig) (*PasswordResetConfig, error)
}

type PasswordResetConfig struct {
	ID              int64
	UsersID         int64
	Url             string
	ExpirationByUse int
	UpdatedAt       int64
}

func NewPasswordResetConfig(userID int64) *PasswordResetConfig {
	passwordResetConfig := &PasswordResetConfig{UsersID: userID}
	passwordResetConfig.Url = fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(time.Now().Unix(), 10)+strconv.FormatInt(userID, 10))))
	passwordResetConfig.ExpirationByUse = 0
	passwordResetConfig.UpdatedAt = internal.Now().UTC().Unix()
	return passwordResetConfig
}
