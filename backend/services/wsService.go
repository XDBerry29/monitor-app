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
	//fmt.Print("Connected")
}

func (s *WsService) RemoveWs(ws *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.wsList, ws)
}

func (s *WsService) SendAll(message []byte) {
	for ws := range s.wsList {
		if err := ws.WriteMessage(websocket.TextMessage, message); err != nil {
			s.RemoveWs(ws)
			fmt.Print("Client Dissconected!")
		}
	}
}

func (s *WsService) SendOne(message []byte, ws *websocket.Conn) {
	if err := ws.WriteMessage(websocket.TextMessage, message); err != nil {
		s.RemoveWs(ws)
		fmt.Print("Client Dissconected!")
	}
}
