package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type UserInfo struct {
	Username string `json:"username"`
	Account  string `json:"account"`
}

func whoamiBaseURL() string {
	if u := os.Getenv("NIMBLE_AUTH_WHOAMI_URL"); u != "" {
		return u
	}
	return "https://api.webit.live"
}

// ValidateToken calls the whoami endpoint with the given bearer token.
// tokenLabel names the credential in error messages (e.g. "API key", "access
// token"), since the endpoint accepts both an API key and an OAuth access
// token.
func ValidateToken(ctx context.Context, token, tokenLabel string) (*UserInfo, error) {
	reqURL := whoamiBaseURL() + "/api/v1/auth/whoami"

	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to reach authentication server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("%s is invalid or expired", tokenLabel)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authentication server returned status %d", resp.StatusCode)
	}

	var info UserInfo
	if err := decodeBody(resp, &info); err != nil {
		return nil, fmt.Errorf("failed to parse authentication response: %w", err)
	}
	return &info, nil
}

func ValidateAPIKey(ctx context.Context, apiKey string) (*UserInfo, error) {
	return ValidateToken(ctx, apiKey, "API key")
}

func ValidateAccessToken(ctx context.Context, accessToken string) (*UserInfo, error) {
	return ValidateToken(ctx, accessToken, "access token")
}
