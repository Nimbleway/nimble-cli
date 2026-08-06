package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const cliKeyName = "Nimble CLI"

type apiKeyEntry struct {
	GUID        string `json:"guid"`
	Key         string `json:"key"`
	KeyName     string `json:"key_name"`
	AccountName string `json:"account_name"`
}

// fetchOrCreateAPIKey always creates a fresh API key. The list endpoint
// (GET /api/v1/account/api-key) masks key secrets (first 4 chars +
// "**********" + last 4), so an existing key can never be reused from it.
//
// A new key is created before any old CLI key is removed: revoking first
// would leave the account with no usable key if creation then failed. Stale
// CLI keys are cleaned up afterwards so repeated logins don't accumulate
// keys. If the account is already at its key limit, the stale keys are
// removed first and creation is retried once.
func fetchOrCreateAPIKey(ctx context.Context, baseURL, token string) (*apiKeyEntry, error) {
	stale, listErr := listCLIKeys(ctx, baseURL, token)

	key, err := createAPIKey(ctx, baseURL, token)
	if err != nil {
		if !errors.Is(err, errKeyLimitExceeded) {
			return nil, err
		}
		if listErr != nil {
			return nil, fmt.Errorf("account is at its API key limit and existing keys could not be listed: %w", listErr)
		}
		if len(stale) == 0 {
			return nil, err
		}
		if delErr := deleteKeys(ctx, baseURL, token, stale); delErr != nil {
			return nil, fmt.Errorf("account is at its API key limit and stale CLI keys could not be removed: %w", delErr)
		}
		stale = nil
		if key, err = createAPIKey(ctx, baseURL, token); err != nil {
			return nil, err
		}
	}

	// Best effort: the new key is already usable, so a failed cleanup must not
	// fail the login. Skip the key just created in case the server reused its GUID.
	for _, guid := range stale {
		if guid != key.GUID {
			_ = deleteAPIKey(ctx, baseURL, token, guid)
		}
	}

	return key, nil
}

// listCLIKeys returns the GUIDs of keys previously created by this CLI.
func listCLIKeys(ctx context.Context, baseURL, token string) ([]string, error) {
	keys, err := listAPIKeys(ctx, baseURL, token)
	if err != nil {
		return nil, err
	}

	var guids []string
	for _, k := range keys {
		if k.KeyName == cliKeyName && k.GUID != "" {
			guids = append(guids, k.GUID)
		}
	}
	return guids, nil
}

func deleteKeys(ctx context.Context, baseURL, token string, guids []string) error {
	for _, guid := range guids {
		if err := deleteAPIKey(ctx, baseURL, token, guid); err != nil {
			return err
		}
	}
	return nil
}

func listAPIKeys(ctx context.Context, baseURL, token string) ([]apiKeyEntry, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"/api/v1/account/api-key", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	var keys []apiKeyEntry
	if err := doJSON(req, &keys); err != nil {
		return nil, err
	}
	return keys, nil
}

func deleteAPIKey(ctx context.Context, baseURL, token, guid string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", baseURL+"/api/v1/account/api-key/"+guid, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// A key that is already gone is the state we wanted.
	if resp.StatusCode == http.StatusNotFound {
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("DELETE %s returned status %d", req.URL.Path, resp.StatusCode)
	}
	return nil
}

// errKeyLimitExceeded reports that the account cannot hold another API key.
// Proxit answers 403 when ResourceLimits.ApiKeys is reached.
var errKeyLimitExceeded = errors.New("account API key limit reached")

func createAPIKey(ctx context.Context, baseURL, token string) (*apiKeyEntry, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", baseURL+"/api/v1/account/api-key", strings.NewReader(fmt.Sprintf(`{"key_name":%q}`, cliKeyName)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return nil, errKeyLimitExceeded
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("POST %s returned status %d", req.URL.Path, resp.StatusCode)
	}

	var key apiKeyEntry
	if err := decodeBody(resp, &key); err != nil {
		return nil, err
	}
	if key.Key == "" {
		return nil, fmt.Errorf("server returned empty API key")
	}
	if strings.Contains(key.Key, "*") {
		return nil, fmt.Errorf("server returned a masked API key instead of the secret")
	}
	return &key, nil
}
