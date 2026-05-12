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

func ValidateAPIKey(ctx context.Context, apiKey string) (*UserInfo, error) {
	reqURL := whoamiBaseURL() + "/api/v1/auth/whoami"

	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to reach authentication server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("API key is invalid or expired")
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
