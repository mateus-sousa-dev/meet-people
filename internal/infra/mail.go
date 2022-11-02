package infra

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func StartSmtpDialer() (*gomail.Dialer, error) {
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}
	n := gomail.NewDialer(os.Getenv("SMTP_HOST"), smtpPort, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))
	return n, nil
}
