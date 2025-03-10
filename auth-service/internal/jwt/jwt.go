package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

type CustomClaims struct {
	UserClaims
	jwt.RegisteredClaims
}

type UserClaims struct {
	FirstName string
	LastName  string
}

func GenerateNewTokenForUser(userClaims UserClaims, expirationTime time.Duration) (string, error) {
	claims := &CustomClaims{
		UserClaims: userClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			Issuer:    "game-metrics/auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
