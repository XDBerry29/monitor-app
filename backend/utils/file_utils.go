package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func CreateLogFile(directory string) (*os.File, error) {
	// Get the current date
	now := time.Now()
	formattedDate := now.Format("02-01-2006") // Format date as zz-mm-yyyy

	// Create the log file name
	logFileName := fmt.Sprintf("%s.log", formattedDate)

	// Create the full path for the log file
	logFilePath := filepath.Join(directory, logFileName)

	// Ensure the directory exists
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// Create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %v", err)
	}

	return file, nil
}
