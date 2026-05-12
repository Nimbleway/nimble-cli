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
			fmt.Printf("You are already logged in as %s.\n", creds.AccountName)
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
			fmt.Println("Browser login is not yet implemented. Please use option 2 to paste an API key.")
			return cli.Exit("", 1)
		case "2":
			return handleAPIKeyLogin(ctx, scanner)
		default:
			fmt.Printf("Invalid choice: %s\n", choice)
			return cli.Exit("", 1)
		}
	},
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

	fmt.Println("Validating API key...")
	info, err := auth.ValidateAPIKey(ctx, apiKey)
	if err != nil {
		fmt.Printf("Authentication failed: %s\n", err)
		return cli.Exit("", 1)
	}

	creds := &auth.Credentials{
		APIKey:      apiKey,
		Source:      "manual",
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

func maskAPIKey(key string) string {
	if len(key) <= 12 {
		return "****"
	}
	return key[:4] + "..." + key[len(key)-4:]
}

func init() {
	Command.Commands = append(Command.Commands,
		&loginCommand,
		&whoamiCommand,
		&logoutCommand,
	)
}
