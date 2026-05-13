package cmd

import (
	"context"
	"fmt"

	"github.com/Nimbleway/nimble-cli/internal/auth"
	"github.com/urfave/cli/v3"
)

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
