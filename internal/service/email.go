package service

import (
	"fmt"
	"github.com/mth-ribeiro-dev/finance-api-go.git/internal/model"
	"gopkg.in/mail.v2"
)

type EmailService struct {
	smtpHost     string
	smtpPort     int
	smtpUsername string
	smtpPassword string
}

func NewEmailService(host string, port int, username, password string) *EmailService {
	return &EmailService{
		smtpHost:     host,
		smtpPort:     port,
		smtpUsername: username,
		smtpPassword: password,
	}
}

func (s *EmailService) SendEmail(data model.EmailData) error {
	m := mail.NewMessage()
	m.SetHeader("From", "seu_email@example.com")
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/plain", fmt.Sprintf("De: %s\n\n%s", data.Name, data.Message))

	d := mail.NewDialer(s.smtpHost, s.smtpPort, s.smtpUsername, s.smtpPassword)

	return d.DialAndSend(m)
}
