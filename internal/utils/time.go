package utils

import "time"

// GetCurrentTimestamp returns the current time in "15:04:05" format.
func GetCurrentTimestamp() string {
	return time.Now().Local().Format("15:04:05.0000")
}

// GetCurrentTimestampUTC returns the current UTC time in "15:04:05" format.
func GetCurrentTimestampUTC() string {
	return time.Now().UTC().Format("15:04:05")
}
