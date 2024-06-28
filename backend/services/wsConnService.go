package services

import (
	"encoding/json"
	"fmt"
	"log"
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
	client := w.clientService.GetClient(id)
	if client == nil {
		fmt.Println("Client not found")
		w.RemoveWs(id)
		return
	}

	defer w.RemoveWs(id)

	conn := client.GetConnWs()

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

		var filterMsg models.ProcessFilterMsg
		err = json.Unmarshal(message, &filterMsg)
		if err != nil {
			log.Printf("Failed to unmarshal JSON message: %v", err)
			continue
		}

		client.SetMonitoring(filterMsg.Name, filterMsg.Monitoring)

	}

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
		if client == nil {
			fmt.Println("Client not found")
			w.RemoveWs(id)
			continue
		}
		if err := client.GetConnWs().WriteMessage(websocket.TextMessage, sMessage); err != nil {
			w.RemoveWs(id)
			fmt.Print("Client Dissconected!")
		}
	}

}

// SendOne implements WsService.
func (w *wsConnService) SendOne(message models.ConnectionMessage, id string) {
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
	if err := client.GetConnWs().WriteMessage(websocket.TextMessage, sMessage); err != nil {
		w.RemoveWs(id)
		fmt.Print("Client Dissconected!")
	}
}
