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
