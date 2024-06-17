package services

import (
	"fmt"
	"sync"

	"github.com/XDBerry29/monitor-app/models"
	"github.com/XDBerry29/monitor-app/repsitories"
	"github.com/XDBerry29/monitor-app/utils"
)

type LogService struct {
	logRepo repsitories.LogRepository
	mu      sync.Mutex
	//ws
}

func NewLogService(logRepo repsitories.LogRepository) *LogService {
	return &LogService{
		logRepo: logRepo,
	}
}

func (s *LogService) ProccesLog(log *models.Log, sendFlag bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	//here we will sent the log to the websoket it will be writen in the console for now
	if sendFlag {
		fmt.Printf("%s\n", utils.LogToWriteString(log))
	}
	s.logRepo.SaveLog(log)

	return nil
}
