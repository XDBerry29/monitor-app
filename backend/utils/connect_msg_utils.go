package utils

import (
	"fmt"
	"strings"

	"github.com/XDBerry29/monitor-app/models"
)

func CreateConnectionMessageNewConn(input string) (*models.ConnectionMessage, error) {
	parts := strings.SplitN(input, "|", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid connection message format")
	}

	connMsg := &models.ConnectionMessage{
		ProcessName: parts[1],
		Timestamp:   parts[0],
		Monitoring:  true,
		Connected:   true,
	}

	return connMsg, nil
}

func CreateConnectionMessage(name, timestamp string, connected bool, monitoring bool) *models.ConnectionMessage {
	return &models.ConnectionMessage{
		ProcessName: name,
		Timestamp:   timestamp,
		Monitoring:  monitoring,
		Connected:   connected,
	}
}
