package repsitories

import (
	"github.com/XDBerry29/monitor-app/models"
)

type LogRepository[R any] interface {
	SaveLog(log *models.Log) error
	UpdateRepo(Repo R) error
}
