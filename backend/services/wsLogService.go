package services

import (
	"encoding/json"
	"fmt"
	"log"
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
	client := w.clientService.GetClient(id)
	if client == nil {
		fmt.Println("Client not found")
		w.RemoveWs(id)
		return
	}

	defer w.RemoveWs(id)

	conn := client.GetLogWs()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket closed unexpectedly: %v", err)
			} else {
				log.Printf("WebSocket read error: %v", err)
			}
			break
		}

		var filterMsg models.LogFilterMessage
		err = json.Unmarshal(message, &filterMsg)
		if err != nil {
			log.Printf("Failed to unmarshal JSON message: %v", err)
			continue
		}

		client.SetMinSeverity(filterMsg.Severity)

	}

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
		if client == nil {
			fmt.Println("Client not found")
			w.RemoveWs(id)
			continue
		}
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
	client := w.clientService.GetClient(id)
	if client == nil {
		fmt.Println("Client not found")
		w.RemoveWs(id)
		return
	}
	sMessage, err := json.Marshal(message)
	if err != nil {
		return
	}
	if message.Level >= client.GetMinSeverity() && client.AmIMonitoring(message.ProcessName) {
		if err := client.GetLogWs().WriteMessage(websocket.TextMessage, sMessage); err != nil {
			w.RemoveWs(id)
			fmt.Print("Client Dissconected!")
		}
	}
}
