package auth

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/khalidrizki/IT-helpdesk/back-end/auth-service/internal/accountsclient"
)

type Service struct {
	repository     *Repository
	accountsClient *accountsclient.Client
	jwtSecret      []byte
}

func NewService(
	repository *Repository,
	accountsClient *accountsclient.Client,
	jwtSecret string,
) *Service {
	return &Service{
		repository:     repository,
		accountsClient: accountsClient,
		jwtSecret:      []byte(jwtSecret),
	}
}

func (s *Service) RegisterAccount(
	ctx context.Context,
	input RegisterInput,
) error {
	var validationErrors []string

	email := strings.TrimSpace(input.Email)

	if email == "" {
		validationErrors = append(validationErrors, "email is required")
	} else {
		addr, err := mail.ParseAddress(email)
		if err != nil || addr.Address != email {
			validationErrors = append(validationErrors, "invalid email format")
		}
	}

	if strings.TrimSpace(input.Password) == "" {
		validationErrors = append(validationErrors, "password is required")
	}

	if strings.TrimSpace(input.ConfirmPassword) == "" {
		validationErrors = append(validationErrors, "confirm password is required")
	}

	if input.Password != input.ConfirmPassword {
		validationErrors = append(validationErrors, "passwords do not match")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}

	userID := uuid.NewString()

	// minta accounts service membuat profil (pakai s.accountsClient.CreateAccount)
	err = s.accountsClient.CreateAccount(
		ctx,
		accountsclient.CreateAccountInput{
			ID:           userID,
			Name:         input.Name,
			DepartmentID: input.DepartmentID,
		},
	)
	// error handling (fmt "creating account profile: %w")
	if err != nil {
		return fmt.Errorf("creating account profile: %w", err)
	}

	// simpan credential di database auth service (pakai s.repository.CreateCredential)
	err = s.repository.CreateCredential(
		ctx,
		CreateCredentialInput{
			UserID:       userID,
			Email:        email,
			PasswordHash: string(hashedPassword),
		},
	)
	// error handling (fmt "creating credential: %w")
	if err != nil {
		return fmt.Errorf("creating credential: %w", err)
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

func (s *Service) GetUserProfile(
	ctx context.Context,
	id string,
) (UserProfile, error) {
	// Fetch the account details from the repository without the password hash
	account, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
		return UserProfile{}, fmt.Errorf("getting user profile: %w", err)
	}

	if !account.IsActive {
		return UserProfile{}, errors.New("account is inactive")
	}

	profile := UserProfile{
		ID:         account.ID,
		Name:       account.Name,
		Email:      account.Email,
		Department: account.Department,
		Role:       account.Role,
		IsActive:   account.IsActive,
	}

	return profile, nil
}

func (s *Service) GetAllAccounts(ctx context.Context) ([]UserProfile, error) {
	accounts, err := s.repository.GetAllAccounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting all accounts: %w", err)
	}

	profiles := make([]UserProfile, 0, len(accounts))
	for _, account := range accounts {
		profiles = append(profiles, UserProfile{
			ID:         account.ID,
			Name:       account.Name,
			Email:      account.Email,
			Department: account.Department,
			Role:       account.Role,
			IsActive:   account.IsActive,
		})
	}
	return profiles, nil
}
