package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Nimbleway/nimble-cli/internal/auth"
	"github.com/urfave/cli/v3"
)

var loginCommand = cli.Command{
	Name:     "login",
	Usage:    "Authenticate with the Nimble API",
	Category: "AUTH",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		scanner := bufio.NewScanner(os.Stdin)

		creds, err := auth.LoadCredentials()
		if err != nil {
			return fmt.Errorf("failed to load credentials: %w", err)
		}
		if creds != nil {
			displayName := creds.AccountName
			if displayName == "" {
				displayName = creds.Email
			}
			if displayName == "" {
				displayName = "(unknown account)"
			}
			fmt.Printf("You are already logged in as %s.\n", displayName)
			fmt.Print("Re-authenticate? [y/N]: ")
			if !scanner.Scan() {
				if scanErr := scanner.Err(); scanErr != nil {
					return fmt.Errorf("failed to read input: %w", scanErr)
				}
				fmt.Println("Login cancelled.")
				return nil
			}
			if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
				fmt.Println("Login cancelled.")
				return nil
			}
		}

		fmt.Println("? How would you like to authenticate?")
		fmt.Println("  1. Browser (recommended)")
		fmt.Println("  2. Paste an API key")
		fmt.Print("Enter choice [1-2]: ")

		if !scanner.Scan() {
			if scanErr := scanner.Err(); scanErr != nil {
				return fmt.Errorf("failed to read input: %w", scanErr)
			}
			fmt.Println("Login cancelled.")
			return cli.Exit("", 1)
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			return handleBrowserLogin(ctx)
		case "2":
			return handleAPIKeyLogin(ctx, scanner)
		default:
			fmt.Printf("Invalid choice: %s\n", choice)
			return cli.Exit("", 1)
		}
	},
}

func handleBrowserLogin(ctx context.Context) error {
	cfg := auth.DefaultOAuthConfig()
	result, err := auth.RunOAuthFlow(ctx, cfg)
	if err != nil {
		fmt.Printf("Browser login failed: %s\n", err)
		return cli.Exit("", 1)
	}

	creds := &auth.Credentials{
		APIKey:      result.APIKey,
		Source:      "oauth",
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
		AccountName: result.AccountName,
	}

	if err := auth.SaveCredentials(creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	displayName := result.AccountName
	if displayName == "" {
		displayName = "(unknown account)"
	}
	fmt.Printf("Successfully logged in as %s.\n", displayName)
	return nil
}

func handleAPIKeyLogin(ctx context.Context, scanner *bufio.Scanner) error {
	fmt.Print("Enter your API key: ")
	if !scanner.Scan() {
		if scanErr := scanner.Err(); scanErr != nil {
			return fmt.Errorf("failed to read input: %w", scanErr)
		}
		fmt.Println("Login cancelled.")
		return cli.Exit("", 1)
	}
	apiKey := strings.TrimSpace(scanner.Text())

	if apiKey == "" {
		fmt.Println("API key cannot be empty.")
		return cli.Exit("", 1)
	}

	return completeLogin(ctx, apiKey, "manual")
}

func completeLogin(ctx context.Context, apiKey, source string) error {
	fmt.Println("Validating API key...")
	info, err := auth.ValidateAPIKey(ctx, apiKey)
	if err != nil {
		fmt.Printf("Authentication failed: %s\n", err)
		return cli.Exit("", 1)
	}

	creds := &auth.Credentials{
		APIKey:      apiKey,
		Source:      source,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
		Email:       info.Username,
		AccountName: info.Account,
	}

	if err := auth.SaveCredentials(creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	fmt.Printf("Successfully logged in as %s (%s).\n", info.Account, info.Username)
	return nil
}

func init() {
	Command.Commands = append(Command.Commands,
		&loginCommand,
		&whoamiCommand,
		&logoutCommand,
	)
}
