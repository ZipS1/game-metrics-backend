package jwt

import (
	"crypto/ed25519"
	"game-metrics/auth-service/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserClaims
	jwt.RegisteredClaims
}

type UserClaims struct {
	FirstName string
	LastName  string
}

func GenerateNewTokenForUser(user models.User, expirationTime time.Duration, privateKey ed25519.PrivateKey) (string, error) {
	claims := &CustomClaims{
		UserClaims: UserClaims{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			Issuer:    "game-metrics/auth-service",
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
