package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

type OAuthConfig struct {
	BaseURL string
}

func DefaultOAuthConfig() OAuthConfig {
	baseURL := "https://api.nimbleway.com"
	if u := os.Getenv("NIMBLE_AUTH_BASE_URL"); u != "" {
		baseURL = u
	}
	return OAuthConfig{BaseURL: baseURL}
}

type OAuthResult struct {
	APIKey      string
	AccountName string
}

func RunOAuthFlow(ctx context.Context, cfg OAuthConfig) (*OAuthResult, error) {
	meta, err := discoverEndpoints(ctx, cfg.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to discover OAuth endpoints: %w", err)
	}

	verifier, challenge, err := generatePKCE()
	if err != nil {
		return nil, fmt.Errorf("failed to generate PKCE: %w", err)
	}

	state, err := randomString(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate state: %w", err)
	}

	cb, err := newCallbackServer(state)
	if err != nil {
		return nil, fmt.Errorf("failed to start callback server: %w", err)
	}
	defer cb.Close()

	clientID, err := registerClient(ctx, meta.registrationEndpoint, cb.RedirectURI())
	if err != nil {
		return nil, fmt.Errorf("failed to register OAuth client: %w", err)
	}

	authURL := buildAuthorizeURL(meta.authorizationEndpoint, clientID, cb.RedirectURI(), challenge, state)
	cb.Start()

	fmt.Println("Opening browser to authenticate...")
	if err := OpenBrowser(authURL); err != nil {
		return nil, fmt.Errorf("failed to open browser: %w", err)
	}
	fmt.Println("Waiting for authentication (press Ctrl+C to cancel)...")

	code, err := cb.WaitForCode(ctx, 5*time.Minute)
	if err != nil {
		return nil, err
	}

	accessToken, err := exchangeCode(ctx, meta.tokenEndpoint, clientID, code, verifier, cb.RedirectURI())
	if err != nil {
		return nil, fmt.Errorf("failed to exchange authorization code: %w", err)
	}

	entry, err := fetchOrCreateAPIKey(ctx, cfg.BaseURL, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch API key: %w", err)
	}

	return &OAuthResult{
		APIKey:      entry.Key,
		AccountName: entry.AccountName,
	}, nil
}

func generatePKCE() (verifier, challenge string, err error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}
	verifier = base64.RawURLEncoding.EncodeToString(b)
	h := sha256.Sum256([]byte(verifier))
	challenge = base64.RawURLEncoding.EncodeToString(h[:])
	return verifier, challenge, nil
}

func randomString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
