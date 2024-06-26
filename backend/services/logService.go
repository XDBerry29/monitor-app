package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/XDBerry29/monitor-app/models"
	"github.com/XDBerry29/monitor-app/repsitories"
	"github.com/XDBerry29/monitor-app/utils"
)

type LogService struct {
	logRepo      repsitories.LogRepository[*os.File]
	mu           sync.Mutex
	wsService    *WsService
	min_severity int
}

func NewLogService(logRepo repsitories.LogRepository[*os.File], wsServce *WsService) *LogService {
	return &LogService{
		logRepo:      logRepo,
		wsService:    wsServce,
		min_severity: 0,
	}
}

func (s *LogService) ProccesLog(logM *models.Log, sendFlag bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	//here we will sent the log to the websoket it will be writen in the console for now
	if sendFlag && logM.Level >= s.min_severity {
		logData, err := json.Marshal(logM)
		if err != nil {
			return err
		}
		s.wsService.SendAll(logData)

		//fmt.Printf("%s", utils.LogToWriteString(logM))
	}
	if err := s.logRepo.SaveLog(logM); err != nil {
		log.Printf("Failed to write to file: %v", err)
	}

	return nil
}

func (l *LogService) UpdateFileOnNewDay(directory string) error {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 1, 0, now.Location())
	ticker := time.NewTicker(time.Until(midnight))
	defer ticker.Stop()

	for range ticker.C {
		file, err := utils.CreateLogFile(directory)
		if err != nil {
			fmt.Printf("Failed to create new log file: %v\n", err)
			return err
		}
		l.logRepo.UpdateRepo(file)
		ticker.Reset(24 * time.Hour)
	}
	return nil
}
