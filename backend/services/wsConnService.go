package services

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/XDBerry29/monitor-app/models"
	"github.com/gorilla/websocket"
)

type wsConnService struct {
	clientService *ClientService
	wsList        map[string]bool
	mu            sync.Mutex
}

func NewWsConnService(clientService *ClientService) WsService[models.ConnectionMessage] {
	return &wsConnService{
		clientService: clientService,
		wsList:        make(map[string]bool),
	}
}

// AddWs implements WsService.
func (w *wsConnService) AddWs(id string, ws *websocket.Conn) {
	w.clientService.HandleNewClientConnForConn(id, ws)
	w.mu.Lock()
	defer w.mu.Unlock()
	w.wsList[id] = true
}

// Listen implements WsService.
func (w *wsConnService) Listen(id string) {
	panic("unimplemented")
}

// RemoveWs implements WsService.
func (w *wsConnService) RemoveWs(id string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.clientService.RemoveClient(id)
	delete(w.wsList, id)
}

// SendAll implements WsService.
func (w *wsConnService) SendAll(message models.ConnectionMessage) {
	sMessage, err := json.Marshal(message)
	if err != nil {
		return
	}
	for id := range w.wsList {
		client := w.clientService.GetClient(id)
		if err := client.GetConnWs().WriteMessage(websocket.TextMessage, sMessage); err != nil {
			w.RemoveWs(id)
			fmt.Print("Client Dissconected!")
		}
	}

}

// SendOne implements WsService.
func (w *wsConnService) SendOne(message models.ConnectionMessage, id string) {
	panic("unimplemented")
}
