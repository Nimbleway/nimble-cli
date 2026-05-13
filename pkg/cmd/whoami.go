package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Nimbleway/nimble-cli/internal/auth"
	"github.com/urfave/cli/v3"
)

var whoamiCommand = cli.Command{
	Name:     "whoami",
	Usage:    "Show current authentication status",
	Category: "AUTH",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		if isAPIKeyFlag() {
			key := cmd.String("api-key")
			fmt.Printf("Authenticated via: --api-key flag\n")
			fmt.Printf("API key: %s\n", maskAPIKey(key))
			return nil
		}

		creds, err := auth.LoadCredentials()
		if err != nil {
			return fmt.Errorf("failed to load credentials: %w", err)
		}

		if creds != nil {
			fmt.Printf("Authenticated via: stored credential (%s)\n", creds.Source)
			fmt.Printf("API key: %s\n", maskAPIKey(creds.APIKey))
			if creds.AccountName != "" {
				fmt.Printf("Account: %s\n", creds.AccountName)
			}
			if creds.Email != "" {
				fmt.Printf("Email: %s\n", creds.Email)
			}
			if creds.CreatedAt != "" {
				fmt.Printf("Logged in: %s\n", creds.CreatedAt)
			}
			return nil
		}

		if envKey := os.Getenv("NIMBLE_API_KEY"); envKey != "" {
			fmt.Printf("Authenticated via: NIMBLE_API_KEY environment variable\n")
			fmt.Printf("API key: %s\n", maskAPIKey(envKey))
			return nil
		}

		fmt.Println("Not authenticated. Run 'nimble login' to authenticate.")
		return cli.Exit("", 1)
	},
}

func maskAPIKey(key string) string {
	if len(key) <= 12 {
		return "****"
	}
	return key[:4] + "..." + key[len(key)-4:]
}
