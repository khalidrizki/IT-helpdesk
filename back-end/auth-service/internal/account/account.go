package account

import "time"

type RegisterInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	DepartmentID    int    `json:"department_id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Account struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Department   string    `json:"department"`
	PasswordHash string    `json:"password_hash"`
	Role         string    `json:"role"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
