package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
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

type asMetadata struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	RegistrationEndpoint  string `json:"registration_endpoint"`
}

type dcrResponse struct {
	ClientID string `json:"client_id"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

type apiKeyEntry struct {
	Key         string `json:"key"`
	KeyName     string `json:"key_name"`
	AccountName string `json:"account_name"`
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

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", fmt.Errorf("failed to start callback server: %w", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	redirectURI := fmt.Sprintf("http://127.0.0.1:%d/callback", port)

	clientID, err := registerClient(ctx, meta.RegistrationEndpoint, redirectURI)
	if err != nil {
		return "", fmt.Errorf("failed to register OAuth client: %w", err)
	}

	authURL := buildAuthorizeURL(meta.AuthorizationEndpoint, clientID, redirectURI, challenge, state)

	codeCh := make(chan string, 1)
	errCh := make(chan error, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if s := r.URL.Query().Get("state"); s != state {
			errCh <- fmt.Errorf("state mismatch")
			http.Error(w, "State mismatch", http.StatusBadRequest)
			return
		}
		if errParam := r.URL.Query().Get("error"); errParam != "" {
			desc := r.URL.Query().Get("error_description")
			errCh <- fmt.Errorf("authorization denied: %s (%s)", errParam, desc)
			fmt.Fprintf(w, "<html><body><h1>Login failed</h1><p>%s</p></body></html>", desc)
			return
		}
		code := r.URL.Query().Get("code")
		if code == "" {
			errCh <- fmt.Errorf("no authorization code received")
			http.Error(w, "No code received", http.StatusBadRequest)
			return
		}
		codeCh <- code
		fmt.Fprint(w, "<html><body><h1>Login successful!</h1><p>You can close this tab.</p></body></html>")
	})

	srv := &http.Server{Handler: mux}
	go srv.Serve(listener)
	defer srv.Close()

	fmt.Println("Opening browser to authenticate...")
	if err := OpenBrowser(authURL); err != nil {
		return "", fmt.Errorf("failed to open browser: %w", err)
	}
	fmt.Println("Waiting for authentication (press Ctrl+C to cancel)...")

	var code string
	select {
	case code = <-codeCh:
	case err := <-errCh:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(5 * time.Minute):
		return "", fmt.Errorf("authentication timed out after 5 minutes")
	}

	accessToken, err := exchangeCode(ctx, meta.TokenEndpoint, clientID, code, verifier, redirectURI)
	if err != nil {
		return "", fmt.Errorf("failed to exchange authorization code: %w", err)
	}

	apiKey, err := fetchOrCreateAPIKey(ctx, cfg.BaseURL, accessToken)
	if err != nil {
		return "", fmt.Errorf("failed to fetch API key: %w", err)
	}

	return apiKey, nil
}

func discoverEndpoints(ctx context.Context, baseURL string) (*asMetadata, error) {
	reqURL := baseURL + "/.well-known/oauth-authorization-server"
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("metadata endpoint returned status %d", resp.StatusCode)
	}
	var meta asMetadata
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return nil, err
	}
	return &meta, nil
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
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("client registration returned status %d", resp.StatusCode)
	}
	var dcr dcrResponse
	if err := json.NewDecoder(resp.Body).Decode(&dcr); err != nil {
		return "", err
	}
	return dcr.ClientID, nil
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
	data := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"code_verifier": {verifier},
		"client_id":     {clientID},
		"redirect_uri":  {redirectURI},
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token endpoint returned status %d", resp.StatusCode)
	}
	var tok tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return "", err
	}
	return tok.AccessToken, nil
}

func fetchOrCreateAPIKey(ctx context.Context, baseURL, token string) (string, error) {
	listURL := baseURL + "/api/v1/account/api-key"
	req, err := http.NewRequestWithContext(ctx, "GET", listURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API key list returned status %d", resp.StatusCode)
	}

	var keys []apiKeyEntry
	if err := json.NewDecoder(resp.Body).Decode(&keys); err != nil {
		return "", err
	}

	if len(keys) > 0 {
		return keys[0].Key, nil
	}

	return createAPIKey(ctx, baseURL, token)
}

func createAPIKey(ctx context.Context, baseURL, token string) (string, error) {
	createURL := baseURL + "/api/v1/account/api-key"
	body := `{"key_name":"Nimble CLI"}`
	req, err := http.NewRequestWithContext(ctx, "POST", createURL, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("API key create returned status %d", resp.StatusCode)
	}
	var key apiKeyEntry
	if err := json.NewDecoder(resp.Body).Decode(&key); err != nil {
		return "", err
	}
	return key.Key, nil
}
