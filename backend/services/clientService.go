package services

import (
	"github.com/XDBerry29/monitor-app/models"
	"github.com/XDBerry29/monitor-app/repsitories"
	"github.com/gorilla/websocket"
)

type ClientService struct {
	clientRepo *repsitories.ClientRepo
}

func NewClientService(clientRepo *repsitories.ClientRepo) *ClientService {
	return &ClientService{clientRepo: clientRepo}
}

func (s *ClientService) HandleNewClientConnForConn(id string, wsc *websocket.Conn) {
	client := s.clientRepo.GetClient(id)
	if client == nil {
		client = models.NewClient(id)
		s.clientRepo.AddClient(client)
	}

	client.SetConnWs(wsc)

}

func (s *ClientService) HandleNewClientConnForLog(id string, wsl *websocket.Conn) {

	client := s.clientRepo.GetClient(id)
	if client == nil {
		client = models.NewClient(id)
		s.clientRepo.AddClient(client)
	}
	client.SetLogWs(wsl)

}

func (s *ClientService) AddOnFilterToAllClients(name string) {
	clist := s.clientRepo.GetAllClients()
	for _, c := range clist {
		c.SetMonitoring(name, true)
	}
}

func (s *ClientService) GetClient(id string) *models.Client {
	return s.clientRepo.GetClient(id)
}

func (s *ClientService) RemoveClient(id string) {
	s.clientRepo.RemoveClient(id)
}
