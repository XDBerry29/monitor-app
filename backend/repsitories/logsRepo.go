package repsitories

import "github.com/XDBerry29/monitor-app/models"

type LogRepository interface {
	SaveLog(log *models.Log) error
}
