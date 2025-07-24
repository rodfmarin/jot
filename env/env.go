// Package env provides utilities for working with environment variables.
package env

import (
	"errors"
	"os"
)

const ObsidianVaultPathEnv = "OBSIDIAN_VAULT_PATH"

// CheckJotEnv checks the environment variables required for Jot to function correctly.
func CheckJotEnv() error {
	// Calls the other check functions to make sure OK
	if err := checkVaultEnv(); err != nil {
		return err
	}
	return nil
}

// checkVaultEnv checks if the Obsidian Vault environment variable is set.
func checkVaultEnv() error {
	// Checks if Obsidian Vault Env Path is set
	if _, ok := os.LookupEnv(ObsidianVaultPathEnv); !ok {
		return errors.New("OBSIDIAN_VAULT_PATH environment variable not set")
	}
	return nil
}
