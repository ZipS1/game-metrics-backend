package handlers

import (
	"crypto/ed25519"
)

type publicKeyProviderMock func() (ed25519.PublicKey, error)

func (f publicKeyProviderMock) GetPublicKey() (ed25519.PublicKey, error) {
	return f()
}
