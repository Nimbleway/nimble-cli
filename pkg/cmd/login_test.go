package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// runCLI runs the nimble CLI binary with the given args and env overrides,
// returning stdout, stderr, and the exit code. It uses "go run" like the
// existing mocktest infrastructure but supports injecting env vars for
// credential isolation.
func runCLI(t *testing.T, env map[string]string, args ...string) (stdout, stderr string, exitCode int) {
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

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	exitCode = 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("Failed to run CLI: %v", err)
	}

	return outBuf.String(), errBuf.String(), exitCode
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

func TestWhoamiEnvOverridesStored(t *testing.T) {
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

	assert.Equal(t, 0, exitCode, "whoami with env var should exit 0")
	// Should show env var source, not the stored credential
	assert.Contains(t, stdout, "NIMBLE_API_KEY")
	assert.NotContains(t, stdout, "stored-account")
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
