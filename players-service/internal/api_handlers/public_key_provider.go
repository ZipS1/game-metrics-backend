package api_handlers

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PublicKeyProvider struct {
	jwksEndpoint string
}

func (p *PublicKeyProvider) Init(jwks string) {
	p.jwksEndpoint = jwks
}

func (p PublicKeyProvider) GetPublicKey() (ed25519.PublicKey, error) {
	var JwksResponse struct {
		Alg  string `json:"alg"`
		Jwks string `json:"jwks"`
	}

	resp, err := http.Get(p.jwksEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error sending jwks request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("jwks endpoint returned %d status", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading jwks response: %w", err)
	}

	if err := json.Unmarshal(body, &JwksResponse); err != nil {
		return nil, fmt.Errorf("invalid jwks response: %w", err)
	}

	keyBytes, err := base64.StdEncoding.DecodeString(JwksResponse.Jwks)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}

	if len(keyBytes) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("invalid public key length")
	}

	return ed25519.PublicKey(keyBytes), nil
}
