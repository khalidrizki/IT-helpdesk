package account

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

var deptList = map[int]string{
	1:  "business-development",
	2:  "customer-service",
	3:  "finance-accounting",
	4:  "human-resources",
	5:  "it",
	6:  "legal-compliance",
	7:  "logistics",
	8:  "management",
	9:  "marketing",
	10: "operations",
	11: "procurement",
	12: "sales",
}

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) RegisterAccount(
	ctx context.Context,
	reqData RegisterInput,
) error {
	var validationErrors []string

	name := strings.TrimSpace(reqData.Name)
	email := strings.TrimSpace(reqData.Email)

	if name == "" {
		validationErrors = append(validationErrors, "name is required")
	} else if utf8.RuneCountInString(name) < 3 {
		validationErrors = append(validationErrors, "name must contain at least 3 characters")
	}

	if email == "" {
		validationErrors = append(validationErrors, "email is required")
	} else {
		addr, err := mail.ParseAddress(email)
		if err != nil || addr.Address != email {
			validationErrors = append(validationErrors, "invalid email format")
		}
	}

	department, exists := deptList[reqData.DepartmentID]
	if !exists {
		validationErrors = append(validationErrors, "invalid department")
	}

	if strings.TrimSpace(reqData.Password) == "" {
		validationErrors = append(validationErrors, "password is required")
	}

	if strings.TrimSpace(reqData.ConfirmPassword) == "" {
		validationErrors = append(validationErrors, "confirm password is required")
	}

	if reqData.Password != reqData.ConfirmPassword {
		validationErrors = append(validationErrors, "passwords do not match")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(reqData.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}
	err = s.repository.CreateAccount(ctx, CreateAccountInput{
		Name:         name,
		Email:        email,
		Department:   department,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		return fmt.Errorf("creating account: %w", err)
	}

	return nil
}
