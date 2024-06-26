package services

import (
	"encoding/json"
	"sync"

	"github.com/XDBerry29/monitor-app/models"
)

type ConnectionService struct {
	wsService *WsService
	mu        sync.Mutex
}

func NewConnectionService(wsService *WsService) *ConnectionService {
	return &ConnectionService{
		wsService: wsService,
	}
}

func (s *ConnectionService) ProccesConnectionMessage(message *models.ConnectionMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sMessage, err := json.Marshal(message)
	if err != nil {
		return
	}
	s.wsService.SendAll(sMessage)

}
