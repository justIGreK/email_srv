package service

import (
	"context"

	"github.com/justIGreK/emailcheck/go/emailcheck"
)

type EmailServiceServer struct {
	emailcheck.UnimplementedEmailServiceServer
}

func (s *EmailServiceServer) ValidateEmail(ctx context.Context, req *emailcheck.EmailRequest) (*emailcheck.EmailValidationResponse, error) {
	isValid, normalizedEmail, localPart, domain := ValidateEmail(req.Email)

	return &emailcheck.EmailValidationResponse{
		IsValid:         isValid,
		NormalizedEmail: normalizedEmail,
		LocalPart:       localPart,
		Domain:          domain,
	}, nil
}
