package usecase

import (
	"errors"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
)

type UserUseCase struct {
	repo           domain.UserRepository
	mailRepository domain.MailRepository
}

func NewUserUseCase(repo domain.UserRepository, mailRepository domain.MailRepository) *UserUseCase {
	return &UserUseCase{repo: repo, mailRepository: mailRepository}
}

func (u *UserUseCase) CreateUser(userDto domain.UserDto) (*domain.User, error) {
	user := u.repo.FindUserByEmail(userDto.Email)
	if user != nil {
		return nil, errors.New("email already exists")
	}
	user, err := domain.NewUser(userDto)
	if err != nil {
		return nil, err
	}
	user, err = u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	mailSender := &domain.MailSender{
		From:        "no-replu@meetpeople.com",
		To:          user.Email,
		Subject:     "Link de ativação",
		ContentType: "text/plain",
		Body:        "Clique no link para ativar a sua conta: ",
	}

	err = u.mailRepository.SendMail(mailSender)
	if err != nil {
		return nil, err
	}
	return user, nil
}
