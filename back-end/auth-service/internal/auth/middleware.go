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
			cookie, err := r.Cookie("access_token")
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			// 2. Parse and verify JWT
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(
				cookie.Value,
				claims,
				func(token *jwt.Token) (any, error) {
					return secretKey, nil
				},
				jwt.WithValidMethods([]string{"HS256"}),
			)

			// 3. Reject invalid token
			if err != nil || !token.Valid {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			if claims.Subject == "" || claims.ExpiresAt == nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			// 4. Pass request to next handler
			next.ServeHTTP(w, r)
		})
	}
}
