package usecase

import (
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	userRepository domain.UserRepository
}

func NewLoginUseCase(userRepository domain.UserRepository) *LoginUseCase {
	return &LoginUseCase{userRepository: userRepository}
}

func (l *LoginUseCase) Exec(loginDto domain.LoginDto) (string, error) {
	user := l.userRepository.FindUserByEmail(loginDto.Email)
	hash, err := bcrypt.GenerateFromPassword([]byte(" "), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	storedPassword := string(hash)
	if user != nil {
		storedPassword = user.Password
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(loginDto.Password))
	if err != nil {
		return "", err
	}
	return "TOKEN", nil
}
