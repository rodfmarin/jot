// Package utils provides utility functions
package utils

import "time"

// GetObsidianDate returns the current date formatted for Obsidian notes
func GetObsidianDate() string {
	// Get the current date with the format of YYYY-MM-DD
	currentTime := GetDate()
	return currentTime.Format("2006-01-02")
}

// GetDate returns the current date and time
func GetDate() time.Time {
	// Return the date
	currentTime := time.Now()
	return currentTime
}
