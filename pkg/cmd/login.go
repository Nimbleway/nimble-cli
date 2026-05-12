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
		if cmd.IsSet("api-key") {
			key := cmd.String("api-key")
			fmt.Printf("Authenticated via: NIMBLE_API_KEY environment variable\n")
			fmt.Printf("API key: %s\n", maskAPIKey(key))
			return nil
		}

		creds, err := auth.LoadCredentials()
		if err != nil {
			return fmt.Errorf("failed to load credentials: %w", err)
		}

		if creds == nil {
			fmt.Println("Not authenticated. Run 'nimble login' to authenticate.")
			os.Exit(1)
		}

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
	},
}

var logoutCommand = cli.Command{
	Name:     "logout",
	Usage:    "Remove stored credentials",
	Category: "AUTH",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		creds, err := auth.LoadCredentials()
		if err != nil {
			return fmt.Errorf("failed to load credentials: %w", err)
		}

		if creds == nil {
			fmt.Println("Not currently logged in.")
			return nil
		}

		if err := auth.DeleteCredentials(); err != nil {
			return fmt.Errorf("failed to delete credentials: %w", err)
		}

		fmt.Println("Successfully logged out.")
		return nil
	},
}

func maskAPIKey(key string) string {
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "..." + key[len(key)-4:]
}

func init() {
	Command.Commands = append(Command.Commands,
		&whoamiCommand,
		&logoutCommand,
	)
}
