package account

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"
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
	jwtSecret  []byte
}

func NewService(repository *Repository, jwtSecret string) *Service {
	return &Service{
		repository: repository,
		jwtSecret:  []byte(jwtSecret),
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

func (s *Service) LoginAccount(
	ctx context.Context,
	reqData LoginInput,
) (string, error) {
	email := strings.TrimSpace(reqData.Email)
	password := strings.TrimSpace(reqData.Password)

	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	account, err := s.repository.GetAccountByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("invalid email or password")
		}

		return "", fmt.Errorf("getting account by email: %w", err)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(account.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := generateJWTToken(account.ID, account.Role, s.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("generating JWT token: %w", err)
	}

	return token, nil
}
