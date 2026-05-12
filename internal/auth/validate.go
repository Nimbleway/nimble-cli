package auth

import (
	"encoding/json"
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

func ValidateAPIKey(apiKey string) (*UserInfo, error) {
	url := whoamiBaseURL() + "/api/v1/auth/whoami"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
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
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("failed to parse authentication response: %w", err)
	}
	return &info, nil
}
