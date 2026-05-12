package auth

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
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

func RunOAuthFlow(ctx context.Context, cfg OAuthConfig) (string, error) {
	meta, err := discoverEndpoints(ctx, cfg.BaseURL)
	if err != nil {
		return "", fmt.Errorf("failed to discover OAuth endpoints: %w", err)
	}

	verifier, challenge, err := generatePKCE()
	if err != nil {
		return "", fmt.Errorf("failed to generate PKCE: %w", err)
	}

	state, err := randomString(32)
	if err != nil {
		return "", fmt.Errorf("failed to generate state: %w", err)
	}

	cb, err := newCallbackServer(state)
	if err != nil {
		return "", fmt.Errorf("failed to start callback server: %w", err)
	}
	defer cb.Close()

	clientID, err := registerClient(ctx, meta.registrationEndpoint, cb.RedirectURI())
	if err != nil {
		return "", fmt.Errorf("failed to register OAuth client: %w", err)
	}

	authURL := buildAuthorizeURL(meta.authorizationEndpoint, clientID, cb.RedirectURI(), challenge, state)
	cb.Start()

	fmt.Println("Opening browser to authenticate...")
	if err := OpenBrowser(authURL); err != nil {
		return "", fmt.Errorf("failed to open browser: %w", err)
	}
	fmt.Println("Waiting for authentication (press Ctrl+C to cancel)...")

	code, err := cb.WaitForCode(ctx, 5*time.Minute)
	if err != nil {
		return "", err
	}

	accessToken, err := exchangeCode(ctx, meta.tokenEndpoint, clientID, code, verifier, cb.RedirectURI())
	if err != nil {
		return "", fmt.Errorf("failed to exchange authorization code: %w", err)
	}

	apiKey, err := fetchOrCreateAPIKey(ctx, cfg.BaseURL, accessToken)
	if err != nil {
		return "", fmt.Errorf("failed to fetch API key: %w", err)
	}

	return apiKey, nil
}

// OAuth endpoint discovery

type oauthEndpoints struct {
	authorizationEndpoint string
	tokenEndpoint         string
	registrationEndpoint  string
}

func discoverEndpoints(ctx context.Context, baseURL string) (*oauthEndpoints, error) {
	reqURL := baseURL + "/.well-known/oauth-authorization-server"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	var raw struct {
		AuthorizationEndpoint string `json:"authorization_endpoint"`
		TokenEndpoint         string `json:"token_endpoint"`
		RegistrationEndpoint  string `json:"registration_endpoint"`
	}
	if err := doJSON(req, &raw); err != nil {
		return nil, err
	}

	return &oauthEndpoints{
		authorizationEndpoint: raw.AuthorizationEndpoint,
		tokenEndpoint:         raw.TokenEndpoint,
		registrationEndpoint:  raw.RegistrationEndpoint,
	}, nil
}

// PKCE (RFC 7636)

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

// OAuth HTTP operations

func registerClient(ctx context.Context, endpoint, redirectURI string) (string, error) {
	body := map[string]interface{}{
		"client_name":                "Nimble CLI",
		"redirect_uris":             []string{redirectURI},
		"grant_types":               []string{"authorization_code"},
		"response_types":            []string{"code"},
		"token_endpoint_auth_method": "none",
	}
	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	var resp struct {
		ClientID string `json:"client_id"`
	}
	if err := doJSON(req, &resp); err != nil {
		return "", err
	}
	return resp.ClientID, nil
}

func buildAuthorizeURL(endpoint, clientID, redirectURI, challenge, state string) string {
	v := url.Values{
		"client_id":             {clientID},
		"redirect_uri":          {redirectURI},
		"response_type":         {"code"},
		"code_challenge":        {challenge},
		"code_challenge_method": {"S256"},
		"state":                 {state},
		"scope":                 {"openid offline_access nimble:read nimble:write"},
	}
	return endpoint + "?" + v.Encode()
}

func exchangeCode(ctx context.Context, endpoint, clientID, code, verifier, redirectURI string) (string, error) {
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"code_verifier": {verifier},
		"client_id":     {clientID},
		"redirect_uri":  {redirectURI},
	}

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var resp struct {
		AccessToken string `json:"access_token"`
	}
	if err := doJSON(req, &resp); err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}

// API key retrieval

type apiKeyEntry struct {
	Key         string `json:"key"`
	KeyName     string `json:"key_name"`
	AccountName string `json:"account_name"`
}

func fetchOrCreateAPIKey(ctx context.Context, baseURL, token string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/api/v1/account/api-key", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	var keys []apiKeyEntry
	if err := doJSON(req, &keys); err != nil {
		return "", err
	}

	if len(keys) > 0 {
		if keys[0].Key == "" {
			return "", fmt.Errorf("server returned empty API key")
		}
		return keys[0].Key, nil
	}

	return createAPIKey(ctx, baseURL, token)
}

func createAPIKey(ctx context.Context, baseURL, token string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", baseURL+"/api/v1/account/api-key", strings.NewReader(`{"key_name":"Nimble CLI"}`))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	var key apiKeyEntry
	if err := doJSON(req, &key); err != nil {
		return "", err
	}
	if key.Key == "" {
		return "", fmt.Errorf("server returned empty API key")
	}
	return key.Key, nil
}

// doJSON executes an HTTP request and decodes the JSON response into dest.
// Returns an error on non-2xx status codes.
func doJSON(req *http.Request, dest interface{}) error {
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%s %s returned status %d", req.Method, req.URL.Path, resp.StatusCode)
	}

	return json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(dest)
}
