package services

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type WsService struct {
	wsList map[*websocket.Conn]bool
	mu     sync.Mutex
}

func NewWsService() *WsService {
	return &WsService{
		wsList: make(map[*websocket.Conn]bool),
	}
}

func (s *WsService) AddWs(ws *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.wsList[ws] = true
	fmt.Print("Connected")
}

func (s *WsService) RemoveWs(ws *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.wsList, ws)
}

func (s *WsService) SendAll(message []byte) {
	for ws := range s.wsList {
		ws.WriteMessage(websocket.TextMessage, message)
		//iff err fuck conection // error handeling later
	}
}
