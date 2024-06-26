package models

type ConnectionMessage struct {
	ProcessName string `json:"name"`
	Timestamp   string `json:"time"`
	Connected   bool   `json:"connected"`
	Monitoring  bool   `json:"monitoring"`
}
