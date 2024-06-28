package services

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/XDBerry29/monitor-app/models"
	"github.com/gorilla/websocket"
)

type wsLogService struct {
	clientService *ClientService
	wsList        map[string]bool
	mu            sync.Mutex
}

func NewWsLogService(clientService *ClientService) WsService[models.Log] {
	return &wsLogService{
		clientService: clientService,
		wsList:        make(map[string]bool),
	}
}

// AddWs implements WsService.
func (w *wsLogService) AddWs(id string, ws *websocket.Conn) {
	w.clientService.HandleNewClientConnForLog(id, ws)
	w.mu.Lock()
	defer w.mu.Unlock()
	w.wsList[id] = true
}

// Listen implements WsService.
func (w *wsLogService) Listen(id string) {
	panic("unimplemented")
}

// RemoveWs implements WsService.
func (w *wsLogService) RemoveWs(id string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.clientService.RemoveClient(id)
	delete(w.wsList, id)
}

// SendAll implements WsService.
func (w *wsLogService) SendAll(message models.Log) {
	sMessage, err := json.Marshal(message)
	if err != nil {
		return
	}
	for id := range w.wsList {
		client := w.clientService.GetClient(id)
		if message.Level >= client.GetMinSeverity() && client.AmIMonitoring(message.ProcessName) {
			if err := client.GetLogWs().WriteMessage(websocket.TextMessage, sMessage); err != nil {
				w.RemoveWs(id)
				fmt.Print("Client Dissconected!")
			}
		}
	}
}

// SendOne implements WsService.
func (w *wsLogService) SendOne(message models.Log, id string) {
	panic("unimplemented")
}
