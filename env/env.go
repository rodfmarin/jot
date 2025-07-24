// Package env provides utilities for working with environment variables.
package env

import (
	"errors"
	"os"
)

const ObsidianVaultPathEnv = "OBSIDIAN_VAULT_PATH"

func CheckJotEnv() error {
	// Calls the other check functions to make sure OK
	if err := checkVaultEnv(); err != nil {
		return err
	}
	return nil
}

func checkVaultEnv() error {
	// Checks if Obsidian Vault Env Path is set
	if _, ok := os.LookupEnv(ObsidianVaultPathEnv); !ok {
		return errors.New("OBSIDIAN_VAULT_PATH environment variable not set")
	}
	return nil
}
