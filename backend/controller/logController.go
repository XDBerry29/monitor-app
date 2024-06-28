package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/XDBerry29/monitor-app/models"
	"github.com/XDBerry29/monitor-app/services"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type LogController struct {
	wsService     services.WsService[models.Log]
	clientService services.ClientService
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewLogController(wsService services.WsService[models.Log], clientService services.ClientService) *LogController {
	return &LogController{
		wsService:     wsService,
		clientService: clientService,
	}
}

func (c *LogController) HandleWs(context echo.Context) error {
	ws, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		return err
	}

	_, client_message, err := ws.ReadMessage()
	if err != nil {
		log.Printf("read error: %v", err)
		return err
	}

	var initMsg map[string]string
	err = json.Unmarshal(client_message, &initMsg)
	if err != nil {
		log.Printf("unmarshal error: %v", err)
		return err
	}

	clientId, ok := initMsg["clientId"]
	if !ok {
		log.Printf("clientId not provided")
		return nil
	}

	c.wsService.AddWs(clientId, ws)

	return nil
}
