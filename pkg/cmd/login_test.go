package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

	cmd.Env = os.Environ()
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
