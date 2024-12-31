package validation

import (
	"fmt"
	"net/mail"
	"strings"
	"unicode"
)

func ValidateEmail(email string) error {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email address: %v", err)
	}
	if addr.Address != email {
		return fmt.Errorf("email address contains invalid characters")
	}
	return nil
}

func SanitizeString(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) && !unicode.IsControl(r) {
			return r
		}
		return -1
	}, input)
}

func ValidatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("port number must be between 1 and 65535")
	}
	return nil
}

func ValidateFilePath(path string) error {
	if strings.Contains(path, "../") || strings.Contains(path, "..\\") {
		return fmt.Errorf("path traversal not allowed")
	}
	return nil
}