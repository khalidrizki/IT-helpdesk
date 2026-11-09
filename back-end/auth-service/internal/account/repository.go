package account

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

type CreateAccountInput struct {
	Name         string
	Email        string
	Department   string
	PasswordHash string
}

func (r *Repository) CreateAccount(
	ctx context.Context,
	input CreateAccountInput,
) error {
	query := `
        INSERT INTO users (
            name,
            email,
            department,
            password_hash
        )
        VALUES ($1, $2, $3, $4)
    `

	_, err := r.db.Exec(
		ctx,
		query,
		input.Name,
		input.Email,
		input.Department,
		input.PasswordHash,
	)

	return err
}

func (r *Repository) GetAccountByEmail(
	ctx context.Context,
	email string,
) (Account, error) {
	var account Account

	query := `
		SELECT id, name, email, department, password_hash, role, is_active
		FROM users
		WHERE lower(email) = lower($1)
	`

	err := r.db.QueryRow(ctx, query, email).Scan(
		&account.ID,
		&account.Name,
		&account.Department,
		&account.PasswordHash,
		&account.Role,
		&account.IsActive,
	)

	if err != nil {
		return Account{}, err
	}

	return account, nil
}
