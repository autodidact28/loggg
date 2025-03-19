package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Log levels
const (
	INFO  = "INFO"
	WARN  = "WARN"
	DEBUG = "DEBUG"
	ERROR = "ERROR"
)

// Get log file path for Linux
func getLogFilePath() string {
	return "/var/log/app/app.log"
}

// Ensure log directory exists
func ensureLogDirectory(logFilePath string) error {
	logDir := logFilePath[:len(logFilePath)-len("app.log")] // Extract directory path
	return os.MkdirAll(logDir, 0755) // Standard Linux permissions
}

// Log message to both file and terminal
func logMessage(level, message string) {
	logMsg := fmt.Sprintf("[%s] %s: %s", level, time.Now().Format(time.RFC3339), message)
	fmt.Println(logMsg) // Print to terminal
	log.Println(logMsg) // Write to log file
}

func main() {
	logFilePath := getLogFilePath()

	// Ensure log directory exists
	if err := ensureLogDirectory(logFilePath); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Open or create log file
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Set log output to both file and terminal
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	fmt.Println("✅ Golang Logging Service Running on Linux...")

	// Log messages every 5 seconds
	for {
		logMessage(INFO, "Application is running")
		logMessage(WARN, "This is a warning message")
		logMessage(DEBUG, "Debugging application state")
		logMessage(ERROR, "An error occurred in the application")
		time.Sleep(5 * time.Second)
	}
}

