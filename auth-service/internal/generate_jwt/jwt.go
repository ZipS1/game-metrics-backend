package generate_jwt

import (
	"crypto/ed25519"
	"game-metrics/auth-service/internal/models"
	lib_jwt "game-metrics/libs/jwt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewTokenForUser(user models.User, expirationTime time.Duration, privateKey ed25519.PrivateKey) (string, error) {
	claims := &lib_jwt.CustomClaims{
		UserClaims: lib_jwt.UserClaims{
			Email:     user.Email,
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
