package models

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)

type Log struct {
	Level       int    `json:"severity"`
	Timestamp   string `json:"time"`
	ProcessName string `json:"process"` //unix
	Message     string `json:"message"`
}

func (l *Log) LevelToString() string {
	switch l.Level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}
