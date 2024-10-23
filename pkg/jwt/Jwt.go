package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Payload struct {
	ID          uint     `json:"id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}
type Claims struct {
	jwt.RegisteredClaims
	Payload Payload `json:"payload"`
}

func GenerateJWT(payload Payload, expTime time.Duration) (string, error) {
	if expTime == 0 {
		expTime = 60 * 24 // Default token expiration time: 24 hours
	}
	// expirationTime := time.Now().Add(expTime)
	expirationTime := time.Now().Add(expTime * time.Minute)
	claims := &Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte(os.Getenv("JWT_SECRET"))
	return token.SignedString(jwtKey)
}
