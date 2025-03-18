package config

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func loadJwtKeys(jwtConfig *JwtTokenConfig) error {
	privData, err := os.ReadFile(jwtConfig.PrivateKeyPemFilepath)
	if err != nil {
		return fmt.Errorf("error reading private key file: %w", err)
	}
	privBlock, _ := pem.Decode(privData)
	if privBlock == nil {
		return errors.New("failed to decode PEM block from private key file")
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(privBlock.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse ED25519 private key: %w", err)
	}
	ed25519Priv, ok := parsedKey.(ed25519.PrivateKey)
	if !ok {
		return errors.New("provided private key is not an ED25519 key")
	}
	jwtConfig.Ed25519PrivateKey = ed25519Priv

	pubData, err := os.ReadFile(jwtConfig.PublicKeyPemFilepath)
	if err != nil {
		return fmt.Errorf("error reading public key file: %w", err)
	}
	pubBlock, _ := pem.Decode(pubData)
	if pubBlock == nil {
		return errors.New("failed to decode PEM block from public key file")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse ED25519 public key: %w", err)
	}
	ed25519Pub, ok := pubInterface.(ed25519.PublicKey)
	if !ok {
		return errors.New("provided public key is not an ED25519 key")
	}
	jwtConfig.Ed25519PublicKey = ed25519Pub

	return nil
}
