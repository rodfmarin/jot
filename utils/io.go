package utils

import (
	"errors"
	"log"
	"os"

	"jot/env"
)

// NoteExtension TODO: Make extension configurable
const NoteExtension = ".md"

type DailyNote struct {
	// Path to the daily note file
	Path string
	// Date of the daily note
	Date string
}

func (f *DailyNote) Read() (string, error) {
	// Read the daily note file and return its content
	// If the file does not exist, return an error
	if f.Path == "" {
		return "", errors.New("daily note file path is empty")
	}
	if f.Date == "" {
		return "", errors.New("daily note date is empty")
	}
	content, err := os.ReadFile(f.Path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (f *DailyNote) Write(content string) error {
	// Write content to the daily note file
	// If the file does not exist, create it
	if f.Path == "" {
		return errors.New("daily note file path is empty")
	}
	if f.Date == "" {
		return errors.New("daily note date is empty")
	}
	// We're setting the perms here to 0644, which means the file is readable and writable by the owner,
	// TODO: Make this configurable
	fh, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open daily note file: %v", err)
	}
	defer func(fh *os.File) {
		err := fh.Close()
		if err != nil {
			log.Fatalf("Failed to close daily note file: %v", err)
		}
	}(fh)

	// Prefix a newline to the content because we don't want to append to the last part of the content.
	//TODO: Make this better
	if _, err := fh.WriteString("\n" + content); err != nil {
		log.Fatalf("Failed to write to daily note file: %v", err)
	}
	return nil

}

func NewDailyNote(date string) (*DailyNote, error) {
	// Create a new DailyNote instance with the given date
	if date == "" {
		return nil, errors.New("date cannot be empty")
	}

	// Construct the file path for the daily note
	vaultPath := os.Getenv(env.ObsidianVaultPathEnv)
	if vaultPath == "" {
		return nil, errors.New("OBSIDIAN_VAULT_PATH environment variable not set")
	}

	filePath := vaultPath + "/" + date + NoteExtension // Assuming daily notes are stored as markdown files
	return &DailyNote{
		Path: filePath,
		Date: date,
	}, nil
}
