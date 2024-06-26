package utils

import (
	"encoding/json"

	"github.com/XDBerry29/monitor-app/models"
)

func CreateConnectionMessageNewConn(input []byte) (*models.ConnectionMessage, error) {
	var connMsg models.ConnectionMessage

	if err := json.Unmarshal(input, &connMsg); err != nil {
		return nil, err
	}

	connMsg.Connected = true
	connMsg.Monitoring = true
	return &connMsg, nil
}

func CreateConnectionMessage(name, timestamp string, connected bool, monitoring bool) *models.ConnectionMessage {
	return &models.ConnectionMessage{
		ProcessName: name,
		Timestamp:   timestamp,
		Monitoring:  monitoring,
		Connected:   connected,
	}
}
