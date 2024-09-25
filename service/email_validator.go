package service

import (
	"net/mail"
	"strings"
)

func ValidateEmail(email string) (bool, string, string, string) {
	normalizedEmail := strings.TrimSpace(strings.ToLower(email))
	
	// Validate email format using the net/mail package
	_, err := mail.ParseAddress(normalizedEmail)
	if err != nil {
		return false, "", "", ""
	}

	// Split email into local part and domain
	parts := strings.Split(normalizedEmail, "@")
	localPart := parts[0]
	domain := parts[1]

	return true, normalizedEmail, localPart, domain
}
