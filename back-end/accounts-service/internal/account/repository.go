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

func (r *Repository) CreateAccount(
	ctx context.Context,
	input CreateAccountInput,
) error {
	query := `
		INSERT INTO accounts (
			id,
			name, 
			department_id
		)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(
		ctx,
		query,
		input.ID,
		input.Name,
		input.DepartmentID,
	)
	return err
}
