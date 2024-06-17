package utils

import (
	"fmt"
	"strings"

	"github.com/XDBerry29/monitor-app/models"
)

func CreateLog(input string) (*models.Log, error) {
	parts := strings.SplitN(input, "|", 4)
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid log format")
	}

	log := &models.Log{
		Level:       parts[0],
		Timestamp:   parts[1],
		ProcessName: parts[2],
		Message:     parts[3],
	}

	return log, nil
}

func LogToWriteString(log *models.Log) string {
	return fmt.Sprintf("[%s|%s|%s] %s", log.Level, log.Timestamp, log.ProcessName, log.Message)
}
