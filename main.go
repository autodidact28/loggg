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

// Service names
const (
	AUTH_SERVICE  = "AuthService"
	DB_SERVICE    = "DatabaseService"
	API_SERVICE   = "APIService"
)

// Get log file path for user directory
func getLogFilePath() string {
	return os.ExpandEnv("$HOME/Tools/my/go-log-service/app.log")
}

// Ensure log directory exists
func ensureLogDirectory(logFilePath string) error {
	logDir := logFilePath[:len(logFilePath)-len("app.log")] // Extract directory path
	return os.MkdirAll(logDir, 0755) // Standard Linux permissions
}

// Log message to both file and terminal
func logMessage(service, level, message string) {
	logMsg := fmt.Sprintf("[%s] [%s] %s: %s", service, level, time.Now().Format(time.RFC3339), message)
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
		logMessage(AUTH_SERVICE, INFO, "Authentication successful")
		logMessage(DB_SERVICE, WARN, "Database connection is slow")
		logMessage(API_SERVICE, DEBUG, "API response time is normal")
		logMessage(AUTH_SERVICE, ERROR, "Failed to authenticate user")
		time.Sleep(5 * time.Second)
	}
}

