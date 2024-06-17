package models

type Log struct {
	Level       string `json:"level"`
	Timestamp   string `json:"timestamp"`
	ProcessName string `json:"process_name"` //unix
	Message     string `json:"message"`
}
