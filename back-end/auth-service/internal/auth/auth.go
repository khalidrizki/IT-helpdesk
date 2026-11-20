package auth

type RegisterInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	DepartmentID    int    `json:"department_id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
