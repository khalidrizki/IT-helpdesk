package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func RequireAuth(secretKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// 1. Read cookie

			// 2. Parse and verify JWT

			// 3. Reject invalid token

			// 4. Pass request to next handler

		})
	}
}
