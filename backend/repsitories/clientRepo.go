package repsitories

import (
	"sync"

	"github.com/XDBerry29/monitor-app/models"
)

type ClientRepo struct {
	clientList map[string]*models.Client
	mu         sync.Mutex
}

func NewClientRepo() *ClientRepo {
	return &ClientRepo{
		clientList: make(map[string]*models.Client),
	}
}

func (s *ClientRepo) AddClient(client *models.Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clientList[client.ID] = client
}

func (s *ClientRepo) GetClient(id string) *models.Client {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.clientList[id]
}

func (s *ClientRepo) RemoveClient(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	//here i need to close conn mby ?
	delete(s.clientList, id)
}

func (s *ClientRepo) GetAllClients() map[string]*models.Client {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.clientList
}
