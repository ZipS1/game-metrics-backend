package api_handlers

import (
	"crypto/ed25519"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading jwks response: %w", err)
	}

	if err := json.Unmarshal(body, &JwksResponse); err != nil {
		return nil, fmt.Errorf("invalid jwks response: %w", err)
	}

	return ed25519.PublicKey(JwksResponse.Jwks), nil
}
