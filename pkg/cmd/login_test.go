package cmd

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runCLIWithStdin(t *testing.T, stdin []byte, env map[string]string, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()

	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok, "Could not get current file path")
	project := filepath.Join(filepath.Dir(filename), "..", "..", "cmd", "nimble")

	fullArgs := append([]string{"run", project}, args...)
	cmd := exec.Command("go", fullArgs...)

	// Build a minimal clean environment to prevent parent env bleed.
	// Only copy essential vars so that NIMBLE_* vars from the developer's
	// shell never leak into tests.
	safeVars := []string{"PATH", "HOME", "GOPATH", "GOROOT", "TMPDIR", "GOMODCACHE", "GOCACHE"}
	for _, key := range safeVars {
		if val, ok := os.LookupEnv(key); ok {
			cmd.Env = append(cmd.Env, key+"="+val)
		}
	}
	for k, v := range env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}

	if stdin != nil {
		cmd.Stdin = bytes.NewReader(stdin)
	}

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	exitCode = 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("Failed to run CLI: %v\nstdout: %s\nstderr: %s", err, outBuf.String(), errBuf.String())
	}

	return outBuf.String(), errBuf.String(), exitCode
}

func runCLI(t *testing.T, env map[string]string, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	return runCLIWithStdin(t, nil, env, args...)
}

// writeCredentialsFile creates a credentials.json in the given config dir
// with the specified fields. This simulates a prior "nimble login".
func writeCredentialsFile(t *testing.T, configDir string, creds map[string]string) {
	t.Helper()

	data, err := json.Marshal(creds)
	require.NoError(t, err)

	err = os.MkdirAll(configDir, 0700)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(configDir, "credentials.json"), data, 0600)
	require.NoError(t, err)
}

func TestWhoamiNoAuth(t *testing.T) {
	configDir := t.TempDir()
	stdout, _, exitCode := runCLI(t, map[string]string{
		"NIMBLE_CONFIG_DIR": configDir,
	}, "whoami")

	assert.Equal(t, 1, exitCode, "whoami with no auth should exit 1")
	assert.Contains(t, stdout, "Not authenticated")
}

func TestWhoamiWithStoredCredential(t *testing.T) {
	configDir := t.TempDir()
	writeCredentialsFile(t, configDir, map[string]string{
		"api_key":      "nbl_abc123secretkey456xyz",
		"source":       "manual",
		"created_at":   "2026-05-10T12:00:00Z",
		"email":        "user@example.com",
		"account_name": "test-account",
	})

	stdout, _, exitCode := runCLI(t, map[string]string{
		"NIMBLE_CONFIG_DIR": configDir,
	}, "whoami")

	assert.Equal(t, 0, exitCode, "whoami with stored credential should exit 0")
	assert.Contains(t, stdout, "test-account")
	// API key should be masked (only show first few and last few chars)
	assert.NotContains(t, stdout, "nbl_abc123secretkey456xyz")
	assert.Contains(t, stdout, "nbl_")
}

func TestWhoamiStoredOverridesEnv(t *testing.T) {
	configDir := t.TempDir()
	writeCredentialsFile(t, configDir, map[string]string{
		"api_key":      "nbl_stored_key_value",
		"source":       "manual",
		"created_at":   "2026-05-10T12:00:00Z",
		"account_name": "stored-account",
	})

	stdout, _, exitCode := runCLI(t, map[string]string{
		"NIMBLE_CONFIG_DIR": configDir,
		"NIMBLE_API_KEY":    "nbl_env_override_key",
	}, "whoami")

	assert.Equal(t, 0, exitCode, "whoami with stored creds should exit 0")
	// Stored credential takes priority over env var per ticket
	assert.Contains(t, stdout, "stored credential")
	assert.Contains(t, stdout, "stored-account")
	assert.NotContains(t, stdout, "NIMBLE_API_KEY")
	assert.NotContains(t, stdout, "nbl_env_override_key")
}

func TestWhoamiEnvFallback(t *testing.T) {
	configDir := t.TempDir()

	stdout, _, exitCode := runCLI(t, map[string]string{
		"NIMBLE_CONFIG_DIR": configDir,
		"NIMBLE_API_KEY":    "nbl_env_fallback_key",
	}, "whoami")

	assert.Equal(t, 0, exitCode, "whoami with env var and no stored creds should exit 0")
	assert.Contains(t, stdout, "NIMBLE_API_KEY")
}

func TestLogoutWithCredentials(t *testing.T) {
	configDir := t.TempDir()
	writeCredentialsFile(t, configDir, map[string]string{
		"api_key":    "nbl_some_key",
		"source":     "manual",
		"created_at": "2026-05-10T12:00:00Z",
	})

	stdout, _, exitCode := runCLI(t, map[string]string{
		"NIMBLE_CONFIG_DIR": configDir,
	}, "logout")

	assert.Equal(t, 0, exitCode, "logout should exit 0")
	assert.Contains(t, stdout, "logged out")

	// Verify credentials file is deleted
	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err), "credentials.json should be deleted after logout")
}

func TestLogoutWarnsAboutEnvVar(t *testing.T) {
	configDir := t.TempDir()
	writeCredentialsFile(t, configDir, map[string]string{
		"api_key":    "nbl_some_key",
		"source":     "manual",
		"created_at": "2026-05-10T12:00:00Z",
	})

	stdout, _, exitCode := runCLI(t, map[string]string{
		"NIMBLE_CONFIG_DIR": configDir,
		"NIMBLE_API_KEY":    "nbl_env_key",
	}, "logout")

	assert.Equal(t, 0, exitCode, "logout should exit 0")
	assert.Contains(t, stdout, "logged out")
	assert.Contains(t, stdout, "NIMBLE_API_KEY environment variable is still set")
}

func TestLogoutNoCredentials(t *testing.T) {
	configDir := t.TempDir()

	stdout, _, exitCode := runCLI(t, map[string]string{
		"NIMBLE_CONFIG_DIR": configDir,
	}, "logout")

	assert.Equal(t, 0, exitCode, "logout with no credentials should exit 0")
	assert.Contains(t, stdout, "Not currently logged in")
}

// mockWhoamiServer returns an httptest.Server that validates API keys via the
// /api/v1/auth/whoami endpoint. Keys in validKeys get a 200 with username/account;
// all others get 401.
func mockWhoamiServer(t *testing.T, validKeys map[string]struct{ Username, Account string }) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/auth/whoami" {
			http.NotFound(w, r)
			return
		}
		apiKey := r.Header.Get("Authorization")
		if len(apiKey) > 7 && apiKey[:7] == "Bearer " {
			apiKey = apiKey[7:]
		}
		info, ok := validKeys[apiKey]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error":"unauthorized"}`)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"username":%q,"account":%q}`, info.Username, info.Account)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestLoginAPIKey(t *testing.T) {
	configDir := t.TempDir()
	srv := mockWhoamiServer(t, map[string]struct{ Username, Account string }{
		"nbl_valid_key_12345678": {Username: "user@example.com", Account: "my-account"},
	})

	// Select "2" for paste, then provide the API key
	stdin := []byte("2\nnbl_valid_key_12345678\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
	}, "login")

	assert.Equal(t, 0, exitCode, "login with valid key should exit 0")
	assert.Contains(t, stdout, "Successfully logged in")

	// Verify credentials file was created
	data, err := os.ReadFile(filepath.Join(configDir, "credentials.json"))
	require.NoError(t, err)

	var creds map[string]string
	require.NoError(t, json.Unmarshal(data, &creds))
	assert.Equal(t, "nbl_valid_key_12345678", creds["api_key"])
	assert.Equal(t, "manual", creds["source"])
	assert.Equal(t, "my-account", creds["account_name"])
	assert.Equal(t, "user@example.com", creds["email"])

	// Verify file permissions
	info, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), info.Mode().Perm())
}

func TestLoginAPIKeyInvalid(t *testing.T) {
	configDir := t.TempDir()
	srv := mockWhoamiServer(t, map[string]struct{ Username, Account string }{})

	stdin := []byte("2\nbad-key\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
	}, "login")

	assert.Equal(t, 1, exitCode, "login with invalid key should exit 1")
	assert.Contains(t, stdout, "Authentication failed")

	// Verify no credentials file was created
	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err), "credentials.json should not exist after failed login")
}

func TestLoginAlreadyLoggedIn(t *testing.T) {
	configDir := t.TempDir()
	writeCredentialsFile(t, configDir, map[string]string{
		"api_key":      "nbl_old_key",
		"source":       "manual",
		"created_at":   "2026-05-10T12:00:00Z",
		"account_name": "old-account",
	})

	srv := mockWhoamiServer(t, map[string]struct{ Username, Account string }{
		"nbl_new_key_87654321": {Username: "new@example.com", Account: "new-account"},
	})

	// Confirm re-auth with "y", select paste, provide new key
	stdin := []byte("y\n2\nnbl_new_key_87654321\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
	}, "login")

	assert.Equal(t, 0, exitCode, "re-login should exit 0")
	assert.Contains(t, stdout, "Successfully logged in")

	// Verify credentials were updated
	data, err := os.ReadFile(filepath.Join(configDir, "credentials.json"))
	require.NoError(t, err)

	var creds map[string]string
	require.NoError(t, json.Unmarshal(data, &creds))
	assert.Equal(t, "nbl_new_key_87654321", creds["api_key"])
	assert.Equal(t, "new-account", creds["account_name"])
}

func TestLoginAPIKeyEmpty(t *testing.T) {
	configDir := t.TempDir()

	// Select paste, then provide empty key
	stdin := []byte("2\n\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR": configDir,
	}, "login")

	assert.Equal(t, 1, exitCode, "login with empty key should exit 1")
	assert.Contains(t, stdout, "API key cannot be empty")

	// Verify no credentials file was created
	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err))
}

type mockOAuthServerOptions struct {
	// tamperState, when true, causes /authorize to append "_tampered" to the
	// state value in the redirect, triggering a state-mismatch error path.
	tamperState bool
	// listStatus, when non-zero, is returned by GET on the api-key collection
	// instead of a key list.
	listStatus int
	// deleteStatus, when non-zero, is returned by DELETE on a key instead of 200.
	deleteStatus int
	// createStatus, when non-zero, is returned by the first POST to the api-key
	// collection instead of 201. Later POSTs succeed, which models the
	// key-limit-then-retry path when set to 403.
	createStatus int
	// forbiddenNotKeyLimit, with createStatus 403, returns a 403 that is not a
	// key-limit error, which must not trigger key deletion.
	forbiddenNotKeyLimit bool
	// createdKey overrides the secret returned by POST. Use to model a server
	// that answers with a masked or empty key.
	createdKey string
	// requests, when non-nil, records every "METHOD path" the server handled.
	requests *[]string
}

// maskKey mimics proxit's DBKeyToApiKey obscuring (first 4 + "**********" +
// last 4) without the overlap or out-of-range slice that a short key would
// cause in the server's own implementation.
func maskKey(key string) string {
	if len(key) <= 8 {
		return strings.Repeat("*", len(key))
	}
	return key[:4] + "**********" + key[len(key)-4:]
}

func mockOAuthServer(t *testing.T, apiKey string) *httptest.Server {
	t.Helper()
	return mockOAuthServerWithOptions(t, apiKey, mockOAuthServerOptions{})
}

func mockOAuthServerWithOptions(t *testing.T, apiKey string, opts mockOAuthServerOptions) *httptest.Server {
	t.Helper()

	// Shared PKCE state: /authorize captures the code_challenge so that
	// /oauth/token can verify the code_verifier.
	var mu sync.Mutex
	var storedCodeChallenge string
	createCalls := 0

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if opts.requests != nil {
			mu.Lock()
			*opts.requests = append(*opts.requests, r.Method+" "+r.URL.Path)
			mu.Unlock()
		}
		switch r.URL.Path {
		case "/.well-known/oauth-authorization-server":
			scheme := "http"
			base := scheme + "://" + r.Host
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{
				"authorization_endpoint": %q,
				"token_endpoint": %q,
				"registration_endpoint": %q
			}`, base+"/authorize", base+"/oauth/token", base+"/oauth/register")

		case "/oauth/register":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"client_id":"test_client_id"}`)

		case "/authorize":
			// Capture the PKCE code_challenge for later verification.
			codeChallenge := r.URL.Query().Get("code_challenge")
			mu.Lock()
			storedCodeChallenge = codeChallenge
			mu.Unlock()

			redirectURI := r.URL.Query().Get("redirect_uri")
			state := r.URL.Query().Get("state")
			if opts.tamperState {
				state = state + "_tampered"
			}
			http.Redirect(w, r, redirectURI+"?code=test_auth_code&state="+state, http.StatusFound)

		case "/oauth/token":
			if err := r.ParseForm(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, `{"error":"invalid_request","error_description":"malformed form body"}`)
				return
			}

			// Validate PKCE: compute S256 of the verifier and compare to the
			// stored challenge.
			codeVerifier := r.FormValue("code_verifier")
			mu.Lock()
			challenge := storedCodeChallenge
			mu.Unlock()

			if challenge != "" && codeVerifier != "" {
				h := sha256.Sum256([]byte(codeVerifier))
				computed := base64.RawURLEncoding.EncodeToString(h[:])
				if computed != challenge {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprint(w, `{"error":"invalid_grant","error_description":"PKCE verification failed"}`)
					return
				}
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"access_token":"test_jwt_token"}`)

		case "/api/v1/account/api-key":
			auth := r.Header.Get("Authorization")
			if auth != "Bearer test_jwt_token" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			switch r.Method {
			case http.MethodGet:
				if opts.listStatus != 0 {
					w.WriteHeader(opts.listStatus)
					return
				}
				// The real list endpoint masks key secrets. The CLI must never
				// store this value.
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `[{"guid":"key-guid-1","key":%q,"key_name":"Nimble CLI","account_name":"oauth-account"}]`, maskKey(apiKey))
			case http.MethodPost:
				mu.Lock()
				createCalls++
				first := createCalls == 1
				mu.Unlock()
				if opts.createStatus != 0 && first {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(opts.createStatus)
					// Proxit reports the key limit through the message body;
					// the CLI must only revoke keys for this specific 403.
					if opts.createStatus == http.StatusForbidden && !opts.forbiddenNotKeyLimit {
						fmt.Fprint(w, `{"message":"max api keys limit reached"}`)
					} else {
						fmt.Fprint(w, `{"message":"forbidden"}`)
					}
					return
				}
				created := apiKey
				if opts.createdKey != "" {
					created = opts.createdKey
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				fmt.Fprintf(w, `{"guid":"key-guid-2","key":%q,"key_name":"Nimble CLI","account_name":"oauth-account"}`, created)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}

		case "/api/v1/account/api-key/key-guid-1":
			if r.Method != http.MethodDelete {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if opts.deleteStatus != 0 {
				w.WriteHeader(opts.deleteStatus)
				return
			}
			w.WriteHeader(http.StatusOK)

		case "/api/v1/auth/whoami":
			// Validate that the caller provides the expected API key.
			auth := r.Header.Get("Authorization")
			expectedBearer := "Bearer " + apiKey
			if auth != expectedBearer {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, `{"error":"unauthorized"}`)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"username":"oauth-user@example.com","account":"oauth-account"}`)

		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

// writeBrowserScript writes a fake browser that follows the authorize URL,
// standing in for a human completing the consent page.
func writeBrowserScript(t *testing.T, dir string) string {
	t.Helper()
	path := filepath.Join(dir, "browser.sh")
	require.NoError(t, os.WriteFile(path, []byte("#!/bin/sh\ncurl -sLo /dev/null \"$1\"\n"), 0755))
	return path
}

func TestLoginBrowser(t *testing.T) {
	configDir := t.TempDir()

	var requests []string
	srv := mockOAuthServerWithOptions(t, "nbl_oauth_test_key_12345", mockOAuthServerOptions{
		requests: &requests,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 0, exitCode, "browser login should exit 0; stdout: %s", stdout)
	assert.Contains(t, stdout, "Successfully logged in")
	assert.Contains(t, stdout, "oauth-account")

	data, err := os.ReadFile(filepath.Join(configDir, "credentials.json"))
	require.NoError(t, err)

	var creds map[string]string
	require.NoError(t, json.Unmarshal(data, &creds))
	assert.Equal(t, "nbl_oauth_test_key_12345", creds["api_key"])
	assert.Equal(t, "oauth", creds["source"])
	assert.Equal(t, "oauth-account", creds["account_name"])
	assert.Equal(t, "oauth-user@example.com", creds["email"])

	// The exact order matters: create the new key, prove it works, and only
	// then revoke the old one. Any other order can leave the account with no
	// usable key.
	assert.Equal(t, []string{
		"GET /api/v1/account/api-key",
		"POST /api/v1/account/api-key",
		"GET /api/v1/auth/whoami",
		"DELETE /api/v1/account/api-key/key-guid-1",
	}, keyRequests(requests))
}

// keyRequests filters out the OAuth handshake so only key lifecycle and
// validation calls remain, in the order they were served.
func keyRequests(requests []string) []string {
	var out []string
	for _, r := range requests {
		switch {
		case strings.HasPrefix(r, "GET /api/v1/account/api-key"),
			strings.HasPrefix(r, "POST /api/v1/account/api-key"),
			strings.HasPrefix(r, "DELETE /api/v1/account/api-key"),
			strings.HasPrefix(r, "GET /api/v1/auth/whoami"):
			out = append(out, r)
		}
	}
	return out
}

// TestLoginBrowserMaskedKey guards the bug this flow originally had: storing
// the masked key returned by the list endpoint instead of a real secret.
func TestLoginBrowserMaskedKey(t *testing.T) {
	configDir := t.TempDir()

	var requests []string
	srv := mockOAuthServerWithOptions(t, "nbl_oauth_masked_key_1234", mockOAuthServerOptions{
		createdKey: "nbl_**********1234",
		requests:   &requests,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 1, exitCode, "masked key should fail login; stdout: %s", stdout)
	assert.Contains(t, stdout, "Browser login failed")

	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err), "masked key must not be stored")

	// createAPIKey must reject the masked secret itself. Reaching whoami would
	// mean that check is gone and the test only passes by luck.
	assert.Equal(t, []string{
		"GET /api/v1/account/api-key",
		"POST /api/v1/account/api-key",
	}, keyRequests(requests), "masked key must be rejected before validation")
}

// TestLoginBrowserWhoamiRejects covers a key that the server mints but the
// whoami endpoint refuses. The old key must survive, since the replacement
// turned out to be unusable.
func TestLoginBrowserWhoamiRejects(t *testing.T) {
	configDir := t.TempDir()

	var requests []string
	// whoami only accepts apiKey, so a different created key is rejected.
	srv := mockOAuthServerWithOptions(t, "nbl_oauth_expected_key", mockOAuthServerOptions{
		createdKey: "nbl_oauth_unexpected_key",
		requests:   &requests,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 1, exitCode, "rejected key should fail login; stdout: %s", stdout)
	assert.Contains(t, stdout, "Browser login failed")

	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err), "unvalidated key must not be stored")

	assert.NotContains(t, keyRequests(requests), "DELETE /api/v1/account/api-key/key-guid-1",
		"stale key must not be revoked when validation fails")
}

// TestLoginBrowserSaveFailure makes the credential write fail after validation
// succeeds, proving cleanup runs only after the credential is durable.
func TestLoginBrowserSaveFailure(t *testing.T) {
	tmp := t.TempDir()

	var requests []string
	srv := mockOAuthServerWithOptions(t, "nbl_oauth_save_fail_key", mockOAuthServerOptions{
		requests: &requests,
	})

	// A regular file where the config directory belongs makes MkdirAll fail.
	configDir := filepath.Join(tmp, "blocked")
	require.NoError(t, os.WriteFile(configDir, []byte("not a directory"), 0600))

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, tmp),
	}, "login")

	assert.NotEqual(t, 0, exitCode, "unsaveable credential should fail login; stdout: %s", stdout)
	assert.NotContains(t, keyRequests(requests), "DELETE /api/v1/account/api-key/key-guid-1",
		"stale key must not be revoked when the credential cannot be saved")
}

// TestLoginBrowserForbiddenNotKeyLimit asserts that a 403 which is not a key
// limit error fails without revoking anything.
func TestLoginBrowserForbiddenNotKeyLimit(t *testing.T) {
	configDir := t.TempDir()

	var requests []string
	srv := mockOAuthServerWithOptions(t, "nbl_oauth_forbidden_key", mockOAuthServerOptions{
		createStatus:         http.StatusForbidden,
		forbiddenNotKeyLimit: true,
		requests:             &requests,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 1, exitCode, "non-key-limit 403 should fail login; stdout: %s", stdout)
	assert.Equal(t, []string{
		"GET /api/v1/account/api-key",
		"POST /api/v1/account/api-key",
	}, keyRequests(requests), "a 403 that is not a key limit must not revoke or retry")

	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err))
}

// TestLoginBrowserKeyLimitListFailure covers the key limit hit when the stale
// keys cannot even be listed, so there is nothing safe to revoke.
func TestLoginBrowserKeyLimitListFailure(t *testing.T) {
	configDir := t.TempDir()

	var requests []string
	srv := mockOAuthServerWithOptions(t, "nbl_oauth_limit_list_key", mockOAuthServerOptions{
		createStatus: http.StatusForbidden,
		listStatus:   http.StatusInternalServerError,
		requests:     &requests,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 1, exitCode, "key limit with unusable list should fail login; stdout: %s", stdout)
	assert.NotContains(t, keyRequests(requests), "DELETE /api/v1/account/api-key/key-guid-1",
		"nothing may be revoked when the key list is unknown")

	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err))
}

// TestLoginBrowserKeyLimitDeleteFailure covers the key limit when the stale key
// cannot be removed, so no slot can be freed for a replacement.
func TestLoginBrowserKeyLimitDeleteFailure(t *testing.T) {
	configDir := t.TempDir()

	var requests []string
	srv := mockOAuthServerWithOptions(t, "nbl_oauth_limit_del_key", mockOAuthServerOptions{
		createStatus: http.StatusForbidden,
		deleteStatus: http.StatusInternalServerError,
		requests:     &requests,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 1, exitCode, "key limit with failed cleanup should fail login; stdout: %s", stdout)
	assert.Equal(t, []string{
		"GET /api/v1/account/api-key",
		"POST /api/v1/account/api-key",
		"DELETE /api/v1/account/api-key/key-guid-1",
	}, keyRequests(requests), "creation must not be retried when the slot was never freed")

	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err))
}

// TestLoginBrowserListFailure asserts login still succeeds when stale-key
// cleanup cannot run: the new key is what matters.
func TestLoginBrowserListFailure(t *testing.T) {
	configDir := t.TempDir()

	srv := mockOAuthServerWithOptions(t, "nbl_oauth_list_fail_key", mockOAuthServerOptions{
		listStatus: http.StatusInternalServerError,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 0, exitCode, "list failure should not fail login; stdout: %s", stdout)

	data, err := os.ReadFile(filepath.Join(configDir, "credentials.json"))
	require.NoError(t, err)
	var creds map[string]string
	require.NoError(t, json.Unmarshal(data, &creds))
	assert.Equal(t, "nbl_oauth_list_fail_key", creds["api_key"])
}

// TestLoginBrowserDeleteFailure asserts a failed cleanup of the stale key does
// not discard the credential that was already created and validated.
func TestLoginBrowserDeleteFailure(t *testing.T) {
	configDir := t.TempDir()

	srv := mockOAuthServerWithOptions(t, "nbl_oauth_del_fail_key", mockOAuthServerOptions{
		deleteStatus: http.StatusInternalServerError,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 0, exitCode, "delete failure should not fail login; stdout: %s", stdout)

	data, err := os.ReadFile(filepath.Join(configDir, "credentials.json"))
	require.NoError(t, err)
	var creds map[string]string
	require.NoError(t, json.Unmarshal(data, &creds))
	assert.Equal(t, "nbl_oauth_del_fail_key", creds["api_key"])
}

// TestLoginBrowserKeyLimit covers the 403-at-limit path: stale CLI keys are
// removed and creation is retried once.
func TestLoginBrowserKeyLimit(t *testing.T) {
	configDir := t.TempDir()

	var requests []string
	srv := mockOAuthServerWithOptions(t, "nbl_oauth_limit_key_123", mockOAuthServerOptions{
		createStatus: http.StatusForbidden,
		requests:     &requests,
	})

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         writeBrowserScript(t, configDir),
	}, "login")

	assert.Equal(t, 0, exitCode, "key limit should be recovered by deleting stale keys; stdout: %s", stdout)

	data, err := os.ReadFile(filepath.Join(configDir, "credentials.json"))
	require.NoError(t, err)
	var creds map[string]string
	require.NoError(t, json.Unmarshal(data, &creds))
	assert.Equal(t, "nbl_oauth_limit_key_123", creds["api_key"])

	// Exactly one retry after exactly one revocation, and no second cleanup
	// pass now that the stale key is already gone.
	assert.Equal(t, []string{
		"GET /api/v1/account/api-key",
		"POST /api/v1/account/api-key",
		"DELETE /api/v1/account/api-key/key-guid-1",
		"POST /api/v1/account/api-key",
		"GET /api/v1/auth/whoami",
	}, keyRequests(requests))
}

func TestLoginBrowserInvalid(t *testing.T) {
	configDir := t.TempDir()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/.well-known/oauth-authorization-server":
			w.WriteHeader(http.StatusInternalServerError)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)

	stdin := []byte("1\n")
	stdout, _, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":    configDir,
		"NIMBLE_AUTH_BASE_URL": srv.URL,
		"NIMBLE_BROWSER":       "echo",
	}, "login")

	assert.Equal(t, 1, exitCode, "browser login with failing server should exit 1")
	assert.Contains(t, stdout, "Browser login failed")

	_, err := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(err))
}

func TestLoginBrowserStateMismatch(t *testing.T) {
	configDir := t.TempDir()

	srv := mockOAuthServerWithOptions(t, "nbl_oauth_state_key", mockOAuthServerOptions{
		tamperState: true,
	})

	browserScript := filepath.Join(configDir, "browser.sh")
	err := os.WriteFile(browserScript, []byte("#!/bin/sh\ncurl -sLo /dev/null \"$1\"\n"), 0755)
	require.NoError(t, err)

	stdin := []byte("1\n")
	stdout, stderr, exitCode := runCLIWithStdin(t, stdin, map[string]string{
		"NIMBLE_CONFIG_DIR":      configDir,
		"NIMBLE_AUTH_BASE_URL":   srv.URL,
		"NIMBLE_AUTH_WHOAMI_URL": srv.URL,
		"NIMBLE_BROWSER":         browserScript,
	}, "login")

	combined := stdout + stderr
	assert.NotEqual(t, 0, exitCode, "browser login with state mismatch should exit non-zero; output: %s", combined)
	assert.True(t,
		strings.Contains(combined, "state") || strings.Contains(combined, "mismatch") || strings.Contains(combined, "Browser login failed"),
		"output should mention state mismatch or browser login failure; got: %s", combined,
	)

	_, statErr := os.Stat(filepath.Join(configDir, "credentials.json"))
	assert.True(t, os.IsNotExist(statErr), "credentials.json should not exist after state mismatch")
}

// TestAuthScenario runs sequential subtests that share a single configDir
// to simulate a full login lifecycle. Subtests are intentionally sequential
// and must NOT use t.Parallel() because each step depends on the state left
// by the previous one (e.g., credentials written by 05_login are read by
// 06_whoami_after_login).
func TestAuthScenario(t *testing.T) {
	apiKey := os.Getenv("NIMBLE_TEST_API_KEY")
	if apiKey == "" {
		t.Skip("NIMBLE_TEST_API_KEY not set; set it to run against staging")
	}

	configDir := t.TempDir()
	credFile := filepath.Join(configDir, "credentials.json")

	env := func(extra ...string) map[string]string {
		m := map[string]string{
			"NIMBLE_CONFIG_DIR": configDir,
		}
		for i := 0; i < len(extra); i += 2 {
			m[extra[i]] = extra[i+1]
		}
		return m
	}

	t.Run("01_whoami_no_auth", func(t *testing.T) {
		stdout, _, exitCode := runCLI(t, env(), "whoami")
		assert.Equal(t, 1, exitCode)
		assert.Contains(t, stdout, "Not authenticated")
	})

	t.Run("02_logout_when_not_logged_in", func(t *testing.T) {
		stdout, _, exitCode := runCLI(t, env(), "logout")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "Not currently logged in")
	})

	t.Run("03_login_invalid_key", func(t *testing.T) {
		stdin := []byte("2\nbad-key\n")
		stdout, _, exitCode := runCLIWithStdin(t, stdin, env(), "login")
		assert.Equal(t, 1, exitCode)
		assert.Contains(t, stdout, "Authentication failed")

		_, err := os.Stat(credFile)
		assert.True(t, os.IsNotExist(err), "credentials.json should not exist after failed login")
	})

	t.Run("04_login_empty_key", func(t *testing.T) {
		stdin := []byte("2\n\n")
		stdout, _, exitCode := runCLIWithStdin(t, stdin, env(), "login")
		assert.Equal(t, 1, exitCode)
		assert.Contains(t, stdout, "API key cannot be empty")
	})

	t.Run("05_login", func(t *testing.T) {
		stdin := []byte("2\n" + apiKey + "\n")
		stdout, _, exitCode := runCLIWithStdin(t, stdin, env(), "login")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "Successfully logged in")

		data, err := os.ReadFile(credFile)
		require.NoError(t, err)

		var creds map[string]string
		require.NoError(t, json.Unmarshal(data, &creds))
		assert.Equal(t, apiKey, creds["api_key"])
		assert.Equal(t, "manual", creds["source"])
		assert.NotEmpty(t, creds["account_name"])

		info, err := os.Stat(credFile)
		require.NoError(t, err)
		assert.Equal(t, os.FileMode(0600), info.Mode().Perm())
	})

	t.Run("06_whoami_after_login", func(t *testing.T) {
		stdout, _, exitCode := runCLI(t, env(), "whoami")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "stored credential")
		assert.Contains(t, stdout, "Account:")
		assert.NotContains(t, stdout, apiKey)
	})

	t.Run("07_whoami_flag_overrides_stored", func(t *testing.T) {
		stdout, _, exitCode := runCLI(t, env(), "whoami", "--api-key", "override_key_xxx")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "--api-key flag")
		assert.NotContains(t, stdout, "stored credential")
	})

	t.Run("08_whoami_env_fallback_not_used", func(t *testing.T) {
		stdout, _, exitCode := runCLI(t, env("NIMBLE_API_KEY", "env_key_yyy"), "whoami")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "stored credential")
		assert.NotContains(t, stdout, "NIMBLE_API_KEY")
	})

	t.Run("09_login_reauth_decline", func(t *testing.T) {
		stdin := []byte("n\n")
		stdout, _, exitCode := runCLIWithStdin(t, stdin, env(), "login")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "already logged in")
		assert.Contains(t, stdout, "Login cancelled")

		data, err := os.ReadFile(credFile)
		require.NoError(t, err)
		var creds map[string]string
		require.NoError(t, json.Unmarshal(data, &creds))
		assert.Equal(t, apiKey, creds["api_key"])
	})

	t.Run("10_login_reauth_accept", func(t *testing.T) {
		stdin := []byte("y\n2\n" + apiKey + "\n")
		stdout, _, exitCode := runCLIWithStdin(t, stdin, env(), "login")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "Successfully logged in")

		data, err := os.ReadFile(credFile)
		require.NoError(t, err)
		var creds map[string]string
		require.NoError(t, json.Unmarshal(data, &creds))
		assert.Equal(t, apiKey, creds["api_key"])
	})

	t.Run("11_whoami_after_reauth", func(t *testing.T) {
		stdout, _, exitCode := runCLI(t, env(), "whoami")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "stored credential")
		assert.Contains(t, stdout, "Account:")
	})

	t.Run("12_logout", func(t *testing.T) {
		stdout, _, exitCode := runCLI(t, env(), "logout")
		assert.Equal(t, 0, exitCode)
		assert.Contains(t, stdout, "logged out")

		_, err := os.Stat(credFile)
		assert.True(t, os.IsNotExist(err), "credentials.json should be deleted after logout")
	})

	t.Run("13_whoami_after_logout", func(t *testing.T) {
		stdout, _, exitCode := runCLI(t, env(), "whoami")
		assert.Equal(t, 1, exitCode)
		assert.Contains(t, stdout, "Not authenticated")
	})
}
