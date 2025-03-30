package jwt

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"game-metrics/auth-service/internal/models"
	"game-metrics/auth-service/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserClaims
	jwt.RegisteredClaims
}

type UserClaims struct {
	Email     string
	FirstName string
	LastName  string
}

func GenerateNewTokenForUser(user models.User, expirationTime time.Duration, privateKey ed25519.PrivateKey) (string, error) {
	claims := &CustomClaims{
		UserClaims: UserClaims{
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

func ValidateToken(tokenString string, key ed25519.PublicKey) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodEdDSA.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	if exists, err := repository.UserIdExists(claims.Subject); err != nil {
		return "", err
	} else if !exists {
		return "", errors.New("user does not exists")
	}

	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		return "", errors.New("token has expired")
	}

	return claims.Subject, nil
}
