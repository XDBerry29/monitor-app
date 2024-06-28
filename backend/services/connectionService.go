package services

import (
	"sync"

	"github.com/XDBerry29/monitor-app/models"
)

type ConnectionService struct {
	wsService WsService[models.ConnectionMessage]
	mu        sync.Mutex
}

func NewConnectionService(wsService WsService[models.ConnectionMessage]) *ConnectionService {
	return &ConnectionService{
		wsService: wsService,
	}
}

func (s *ConnectionService) ProccesConnectionMessage(message *models.ConnectionMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.wsService.SendAll(*message)

}
