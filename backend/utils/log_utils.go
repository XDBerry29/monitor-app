package utils

import (
	"encoding/json"
	"fmt"

	"github.com/XDBerry29/monitor-app/models"
)

func CreateLog(input []byte) (*models.Log, error) {
	var log models.Log

	if err := json.Unmarshal(input, &log); err != nil {
		return nil, err
	}
	return &log, nil
}

func LogToWriteString(log *models.Log) string {
	return fmt.Sprintf("[%s|%s|%s] %s\n", log.LevelToString(), log.Timestamp, log.ProcessName, log.Message)
}
