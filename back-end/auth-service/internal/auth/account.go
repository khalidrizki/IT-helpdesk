package auth

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Role       string `json:"role"`
	IsActive   bool   `json:"is_active"`
}
