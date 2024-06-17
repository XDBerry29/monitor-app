package repsitories

import (
	"os"

	"github.com/XDBerry29/monitor-app/models"
)

type Repo interface {
	*os.File
}

type LogRepository[R Repo] interface {
	SaveLog(log *models.Log) error
	UpdateRepo(Repo R) error
}
