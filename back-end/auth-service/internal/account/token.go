package account

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateJWTToken(accountID string, role string, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{
		"sub":  accountID,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
