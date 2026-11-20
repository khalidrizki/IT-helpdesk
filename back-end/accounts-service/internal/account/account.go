package account

import "time"

type CreateAccountInput struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DepartmentID int    `json:"department_id"`
}

type Account struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	DepartmentID string    `json:"department_id"`
	Role         string    `json:"role"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
