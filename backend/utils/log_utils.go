package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/XDBerry29/monitor-app/models"
)

func CreateLog(input string) (*models.Log, error) {
	parts := strings.SplitN(input, "|", 4)
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid log format")
	}

	SeverityLevel, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid severity level: %s", parts[0])
	}

	log := &models.Log{
		Level:       SeverityLevel,
		Timestamp:   parts[1],
		ProcessName: parts[2],
		Message:     parts[3],
	}

	return log, nil
}

func LogToWriteString(log *models.Log) string {
	return fmt.Sprintf("[%s|%s|%s] %s", log.LevelToString(), log.Timestamp, log.ProcessName, log.Message)
}
