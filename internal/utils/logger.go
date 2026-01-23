package utils

import (
	"fmt"
	"io"
	"os"
)

var (
	logFile io.Writer
)

func init() {
	// Open or create result.txt in the root project directory (relative to where app runs)
	f, err := os.OpenFile("scripts/result.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		logFile = os.Stdout
		return
	}
	logFile = io.MultiWriter(os.Stdout, f)
}

// Log formats and prints a message with the current localized timestamp.
func Log(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	fmt.Fprintf(logFile, "[%s] %s\n", GetCurrentTimestamp(), message)
}

// LogRaw formats and prints a message WITHOUT a timestamp.
func LogRaw(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	fmt.Fprintf(logFile, "%s\n", message)
}

// LogError formats and prints an error message.
func LogError(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	fmt.Fprintf(logFile, "[%s] ERROR: %s\n", GetCurrentTimestamp(), message)
}
