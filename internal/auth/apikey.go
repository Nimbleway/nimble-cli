package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// cliKeyPrefix marks a key as CLI-created. It deliberately omits "Nimble":
// every key in the account is a Nimble key, so the word only spends budget
// against maxKeyNameLen.
const cliKeyPrefix = "CLI"

// maxKeyNameLen is the server's limit on the key_name field.
const maxKeyNameLen = 50

// CLIKeyName scopes the key name to the user and machine that created it, so
// two people (or two machines) sharing an account don't revoke each other's
// CLI keys on login. The email domain and FQDN suffix are dropped, and the
// remainder is truncated as needed, to stay within maxKeyNameLen.
func CLIKeyName(username string) string {
	host, err := os.Hostname()
	if err != nil {
		host = ""
	}
	return cliKeyNameFor(username, host)
}

// cliKeyNameFor is CLIKeyName with the hostname injected, so the length and
// truncation rules can be tested without depending on the test machine.
func cliKeyNameFor(username, host string) string {
	if username == "" {
		username = "unknown-user"
	} else if at := strings.IndexByte(username, '@'); at > 0 {
		username = username[:at]
	}

	if host == "" {
		host = "unknown-host"
	} else if dot := strings.IndexByte(host, '.'); dot > 0 {
		host = host[:dot]
	}

	// Trim only as much as the limit demands. Each part has a fair share, but
	// whichever part is under its share lends the surplus to the other, so a
	// short username leaves room for a long hostname and vice versa.
	budget := maxKeyNameLen - len(cliKeyPrefix) - len(" ( @ )")
	fairUser := max(budget*2/3, 1) // favors the username: it identifies the person
	fairHost := max(budget-fairUser, 1)

	if len(username)+len(host) > budget {
		switch {
		case len(username) <= fairUser:
			host = truncate(host, budget-len(username))
		case len(host) <= fairHost:
			username = truncate(username, budget-len(host))
		default:
			username = truncate(username, fairUser)
			host = truncate(host, fairHost)
		}
	}

	return fmt.Sprintf("%s (%s @ %s)", cliKeyPrefix, username, host)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n]
}

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
// would leave the account with no usable key if creation then failed. The
// returned cleanup function removes stale CLI keys so repeated logins don't
// accumulate them, and the caller must not invoke it until the new key has
// been validated and persisted. cleanup is never nil.
//
// If the account is already at its key limit, stale keys are removed up front
// (there is no way to mint a replacement otherwise) and creation is retried
// once.
func fetchOrCreateAPIKey(ctx context.Context, baseURL, token, username string) (key *apiKeyEntry, cleanup func(), err error) {
	noCleanup := func() {}
	keyName := CLIKeyName(username)

	stale, listErr := listCLIKeys(ctx, baseURL, token, keyName)

	key, err = createAPIKey(ctx, baseURL, token, keyName)
	if err != nil {
		if !errors.Is(err, errKeyLimitExceeded) {
			return nil, noCleanup, err
		}
		if listErr != nil {
			return nil, noCleanup, fmt.Errorf("account is at its API key limit and existing keys could not be listed: %w", listErr)
		}
		if len(stale) == 0 {
			return nil, noCleanup, err
		}
		if delErr := deleteKeys(ctx, baseURL, token, stale); delErr != nil {
			return nil, noCleanup, fmt.Errorf("account is at its API key limit and stale CLI keys could not be removed: %w", delErr)
		}
		// Already deleted, so nothing is left for the caller to clean up.
		stale = nil
		if key, err = createAPIKey(ctx, baseURL, token, keyName); err != nil {
			return nil, noCleanup, err
		}
	}

	newGUID := key.GUID
	return key, func() {
		// Best effort: the new key is live and stored by now, so a failed
		// cleanup must not fail the login. Skip the key just created in case
		// the server reused its GUID.
		for _, guid := range stale {
			if guid != newGUID {
				_ = deleteAPIKey(ctx, baseURL, token, guid)
			}
		}
	}, nil
}

// listCLIKeys returns the GUIDs of keys previously created by this CLI for
// the same scoped keyName.
func listCLIKeys(ctx context.Context, baseURL, token, keyName string) ([]string, error) {
	keys, err := listAPIKeys(ctx, baseURL, token)
	if err != nil {
		return nil, err
	}

	var guids []string
	for _, k := range keys {
		if k.KeyName == keyName && k.GUID != "" {
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
// Proxit answers 403 with entities.ErrApiKeysExceededLimit in the message
// field when ResourceLimits.ApiKeys is reached. Other 403s (a read-only
// super-user, a revoked scope) must not be mistaken for it, since only the
// key limit justifies revoking existing keys.
var errKeyLimitExceeded = errors.New("account API key limit reached")

const keyLimitMessage = "api key limit reached"

func isKeyLimitResponse(resp *http.Response) bool {
	var body struct {
		Message string `json:"message"`
	}
	if err := decodeBody(resp, &body); err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(body.Message), keyLimitMessage)
}

func createAPIKey(ctx context.Context, baseURL, token, keyName string) (*apiKeyEntry, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", baseURL+"/api/v1/account/api-key", strings.NewReader(fmt.Sprintf(`{"key_name":%q}`, keyName)))
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

	if resp.StatusCode == http.StatusForbidden && isKeyLimitResponse(resp) {
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
