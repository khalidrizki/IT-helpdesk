package auth

import "context"

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
		&account.Email,
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

func (r *Repository) GetAccountByID(
	ctx context.Context,
	id string,
) (Account, error) {
	var account Account

	query := `
		SELECT id, name, email, department, password_hash, role, is_active
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
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

func (r *Repository) GetAllAccounts(ctx context.Context) ([]Account, error) {

	query := `
		SELECT id, name, email, department, role, is_active
		FROM users
		ORDER BY name ASC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []Account{}
	for rows.Next() {
		var account Account

		err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.Email,
			&account.Department,
			&account.Role,
			&account.IsActive,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
