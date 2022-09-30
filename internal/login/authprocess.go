package login

import (
	"github.com/mateus-sousa-dev/meet-people/app/auth"
	"github.com/mateus-sousa-dev/meet-people/internal/users"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Exec(loginDto LoginDto) (string, error)
}

type loginUseCase struct {
	userRepository users.Repository
}

func NewLoginUseCase(userRepository users.Repository) LoginUseCase {
	return &loginUseCase{userRepository: userRepository}
}

func (l *loginUseCase) Exec(loginDto LoginDto) (string, error) {
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
	token, _ := auth.CreateToken(user.ID)
	return token, nil
}
