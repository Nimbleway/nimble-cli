package auth

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Credentials struct {
	APIKey      string `json:"api_key"`
	Source      string `json:"source"`
	CreatedAt   string `json:"created_at"`
	Email       string `json:"email,omitempty"`
	AccountName string `json:"account_name,omitempty"`
}

func credentialsDir() string {
	if dir := os.Getenv("NIMBLE_CONFIG_DIR"); dir != "" {
		return dir
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".nimble")
}

func CredentialsPath() string {
	return filepath.Join(credentialsDir(), "credentials.json")
}

func LoadCredentials() (*Credentials, error) {
	data, err := os.ReadFile(CredentialsPath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, err
	}
	return &creds, nil
}

func SaveCredentials(creds *Credentials) error {
	dir := credentialsDir()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return err
	}

	// Atomic write: write to temp file, then rename
	tmp := CredentialsPath() + ".tmp"
	if err := os.WriteFile(tmp, data, 0600); err != nil {
		return err
	}
	return os.Rename(tmp, CredentialsPath())
}

func DeleteCredentials() error {
	err := os.Remove(CredentialsPath())
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}
