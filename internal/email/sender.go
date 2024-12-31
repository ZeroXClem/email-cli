package email

import (
	"fmt"
	"net/smtp"
)

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string

func validateEmailConfig(config EmailConfig) error {
	if err := validation.ValidateEmail(config.SMTPUsername); err != nil {
		return fmt.Errorf("invalid SMTP username: %v", err)
	}
	if err := validation.ValidatePort(config.SMTPPort); err != nil {
		return fmt.Errorf("invalid SMTP port: %v", err)
	}
	return nil
}

func createTLSConfig(tlsConfig *config.TLSConfig) (*tls.Config, error) {
	if tlsConfig == nil {
		return &tls.Config{}, nil
	}

	config := &tls.Config{
		InsecureSkipVerify: tlsConfig.SkipVerify,
	}

	if tlsConfig.ServerName != "" {
		config.ServerName = tlsConfig.ServerName
	}

	if tlsConfig.CertificatePath != "" && tlsConfig.PrivateKeyPath != "" {
		cert, err := tls.LoadX509KeyPair(tlsConfig.CertificatePath, tlsConfig.PrivateKeyPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load TLS certificates: %v", err)
		}
		config.Certificates = []tls.Certificate{cert}
	}

	return config, nil
}
}
type Email struct {
	To          []string
	Cc          []string
	Bcc         []string
	From        string
	ReplyTo     string
	Subject     string
func SendEmail(config EmailConfig, email Email) error {
	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost)

	boundary := "boundary-" + time.Now().String()
	headers := make(map[string]string)
	headers["From"] = email.From
	headers["To"] = strings.Join(email.To, ", ")
	if len(email.Cc) > 0 {
		headers["Cc"] = strings.Join(email.Cc, ", ")
	}
	if email.ReplyTo != "" {
		headers["Reply-To"] = email.ReplyTo
	}
	headers["Subject"] = email.Subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "multipart/alternative; boundary=" + boundary

	var message bytes.Buffer

	// Add headers
	for key, value := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	message.WriteString("\r\n")

	// Add plain text body
	message.WriteString("--" + boundary + "\r\n")
	message.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	message.WriteString("\r\n")
	message.WriteString(email.Body)
	message.WriteString("\r\n")

	// Add HTML body if present
	if email.HTMLBody != "" {
		message.WriteString("--" + boundary + "\r\n")
		message.WriteString("Content-Type: text/html; charset=utf-8\r\n")
		message.WriteString("\r\n")
		message.WriteString(email.HTMLBody)
		message.WriteString("\r\n")
	}

	// Add attachments
	for _, attachment := range email.Attachments {
		message.WriteString("\r\n--" + boundary + "\r\n")
		message.WriteString(fmt.Sprintf("Content-Type: %s\r\n", attachment.ContentType))
		message.WriteString("Content-Transfer-Encoding: base64\r\n")
		message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\r\n", attachment.Filename))
		message.WriteString("\r\n")
		b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Data)))
		base64.StdEncoding.Encode(b, attachment.Data)
		message.Write(b)
	}

	message.WriteString("\r\n--" + boundary + "--\r\n")

	// Combine all recipients
	allRecipients := append(append(email.To, email.Cc...), email.Bcc...)

	addr := fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort)
	return smtp.SendMail(addr, auth, email.From, allRecipients, message.Bytes())
}
	message.WriteString(email.Body)
	message.WriteString("\r\n")

	// Add attachments
	for _, attachment := range email.Attachments {
		message.WriteString("\r\n--" + boundary + "\r\n")
		message.WriteString(fmt.Sprintf("Content-Type: %s\r\n", attachment.ContentType))
		message.WriteString("Content-Transfer-Encoding: base64\r\n")
		message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\r\n", attachment.Filename))
		message.WriteString("\r\n")
		b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Data)))
		base64.StdEncoding.Encode(b, attachment.Data)
		message.Write(b)
	}

	message.WriteString("\r\n--" + boundary + "--\r\n")

	addr := fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort)
	return smtp.SendMail(addr, auth, email.From, email.To, message.Bytes())
}
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",
		email.From,
		email.To[0],
		email.Subject,
		email.Body)


func SendTemplatedEmail(config EmailConfig, email Email, templateName string, templateData map[string]interface{}) error {
	tmpl, err := template.LoadTemplate(templateName)
	if err != nil {
		return err
	}

	data := template.TemplateData{
		To:      email.To[0],
		From:    email.From,
		Subject: email.Subject,
		Body:    templateData,
	}

	msg, err := template.ParseTemplate(tmpl, data)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost)
	addr := fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort)
	return smtp.SendMail(addr, auth, email.From, email.To, []byte(msg))
}
	addr := fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort)
	return smtp.SendMail(addr, auth, email.From, email.To, []byte(msg))
}