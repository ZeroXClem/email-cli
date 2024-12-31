package email

import (
	"testing"
	"net/smtp"
)

type mockSMTP struct {
	sent bool
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}

func (m *mockSMTP) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	m.sent = true
	m.auth = a
	m.from = from
	m.to = to
	m.msg = msg
	return nil
}

func TestSendEmail(t *testing.T) {
	mock := &mockSMTP{}

	config := EmailConfig{
		SMTPHost: "smtp.test.com",
		SMTPPort: 587,
		SMTPUsername: "test@test.com",
		SMTPPassword: "password",
	}

	email := Email{
		To: []string{"recipient@test.com"},
		From: "sender@test.com",
		Subject: "Test Email",
		Body: "Test body",
	}

	err := SendEmail(config, email)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !mock.sent {
		t.Error("Expected email to be sent")
	}
}