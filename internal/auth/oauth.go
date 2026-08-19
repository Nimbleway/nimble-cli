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
	// CleanupStaleKeys removes API keys left by earlier CLI logins. It is
	// never nil and must be called only after APIKey has been validated and
	// persisted, so a failure in between cannot revoke the previous key while
	// leaving no usable replacement.
	CleanupStaleKeys func()
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

	// The key name is scoped per user (and per machine) so that two people
	// sharing an account don't revoke each other's CLI keys on login.
	username, err := authenticatedUsername(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	entry, cleanup, err := fetchOrCreateAPIKey(ctx, cfg.BaseURL, accessToken, username)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch API key: %w", err)
	}

	return &OAuthResult{
		APIKey:           entry.Key,
		AccountName:      entry.AccountName,
		CleanupStaleKeys: cleanup,
	}, nil
}

// authenticatedUsername identifies the person logging in, which scopes both the
// key name and the cleanup that follows.
//
// The whoami endpoint is asked first, since it is the authority. When it cannot
// answer, the access token's own username claim is used instead: the name only
// has to be stable and specific to this user, and failing the whole login over
// a call made purely to derive a name would be a worse outcome. If neither
// yields a name, the login stops rather than fall back to a shared name that
// would put another person's key in scope for deletion.
func authenticatedUsername(ctx context.Context, accessToken string) (string, error) {
	info, err := ValidateAccessToken(ctx, accessToken)
	if err == nil && info.Username != "" {
		return info.Username, nil
	}

	if username := tokenUsername(accessToken); username != "" {
		return username, nil
	}

	if err != nil {
		return "", fmt.Errorf("failed to identify authenticated user: %w", err)
	}
	return "", fmt.Errorf("failed to identify authenticated user: no username in the whoami response or the access token")
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
