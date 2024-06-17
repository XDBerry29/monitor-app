package utils

import (
	"fmt"
	"strings"

	"github.com/XDBerry29/monitor-app/models"
)

func CreateConnectionMessage(input string) (*models.ConnectionMessage, error) {
	parts := strings.SplitN(input, "|", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid connection message format")
	}

	connMsg := &models.ConnectionMessage{
		ProcessName: parts[1],
		Timestamp:   parts[0],
	}

	return connMsg, nil
}
