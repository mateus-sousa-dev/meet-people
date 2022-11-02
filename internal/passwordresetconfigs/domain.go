package passwordresetconfigs

import (
	"crypto/md5"
	"fmt"
	"github.com/mateus-sousa-dev/meet-people/internal/global"
	"strconv"
	"time"
)

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
	passwordResetConfig.UpdatedAt = global.Now().UTC().Unix()
	return passwordResetConfig
}

func (p *PasswordResetConfig) IsValidUrl() bool {
	if p.ExpirationByUse != 0 || (global.Now().UTC().Unix()-p.UpdatedAt) > 300 {
		return false
	}
	return true
}
