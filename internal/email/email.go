package email

import (
	"fmt"
	"net/smtp"
)

type Email struct {
	From        string
	To          []string
	Subject     string
	Body        string
	Attachments []string
}

func NewEmail(from string, to []string, subject, body string) *Email {
	return &Email{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (e *Email) Send(host string, port int, username, password string) error {
	// TODO: Implement email sending logic
	return nil
}
