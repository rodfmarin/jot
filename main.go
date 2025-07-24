package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"jot/env"
	"jot/utils"
)

func main() {
	// Check if the environment is set up correctly
	if err := env.CheckJotEnv(); err != nil {
		log.Fatalf("Environment check failed: %v", err)
	}

	// check command line arguments
	args := os.Args[1:] // Exclude the first argument which is the program name
	if err := checkArgs(args); err != nil {
		log.Fatalf("Process args failed: %v", err)
	}

	// Create a flat string from the command line args/strings
	note := strings.Join(args, " ")

	// Create a daily note file
	date := utils.GetObsidianDate()
	dailyNote, err := utils.NewDailyNote(date)
	if err != nil {
		log.Fatalf("Failed to create daily note: %v", err)
	}
	// Write the note to the daily note file
	err = dailyNote.Write(note)
	if err != nil {
		log.Fatalf("Failed to write to daily note: %v", err)
	}
}

func checkArgs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no note provided. please provide a note to write to the daily note file")
	}
	return nil
}
