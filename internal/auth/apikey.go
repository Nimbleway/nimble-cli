package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type apiKeyEntry struct {
	Key         string `json:"key"`
	KeyName     string `json:"key_name"`
	AccountName string `json:"account_name"`
}

func fetchOrCreateAPIKey(ctx context.Context, baseURL, token string) (*apiKeyEntry, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/api/v1/account/api-key", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	var keys []apiKeyEntry
	if err := doJSON(req, &keys); err != nil {
		return nil, err
	}

	if len(keys) > 0 {
		if keys[0].Key == "" {
			return nil, fmt.Errorf("server returned empty API key")
		}
		return &keys[0], nil
	}

	return createAPIKey(ctx, baseURL, token)
}

func createAPIKey(ctx context.Context, baseURL, token string) (*apiKeyEntry, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", baseURL+"/api/v1/account/api-key", strings.NewReader(`{"key_name":"Nimble CLI"}`))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	var key apiKeyEntry
	if err := doJSON(req, &key); err != nil {
		return nil, err
	}
	if key.Key == "" {
		return nil, fmt.Errorf("server returned empty API key")
	}
	return &key, nil
}
