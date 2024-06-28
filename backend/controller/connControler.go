package controller

import (
	"encoding/json"
	"log"

	"github.com/XDBerry29/monitor-app/connections"
	"github.com/XDBerry29/monitor-app/models"
	"github.com/XDBerry29/monitor-app/services"
	"github.com/XDBerry29/monitor-app/utils"
	"github.com/labstack/echo/v4"
)

type ConnectionController struct {
	wsService     services.WsService[models.ConnectionMessage]
	clientService *services.ClientService
	conHub        connections.ConnectionHub
}

func NewConnectionController(wsService services.WsService[models.ConnectionMessage], clientService *services.ClientService, conHub connections.ConnectionHub) *ConnectionController {
	return &ConnectionController{
		wsService:     wsService,
		clientService: clientService,
		conHub:        conHub,
	}
}

// HandleWs handles websocket connections
func (c *ConnectionController) HandleWs(context echo.Context) error {
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

	proccesConnection := c.conHub.GetAllConnections()
	client := c.clientService.GetClient(clientId)

	for _, connection := range proccesConnection {
		message := utils.CreateConnectionMessage(connection.GetName(), "", true)
		client.SetMonitoring(clientId, true)
		c.wsService.SendOne(*message, clientId)
	}

	return nil

}
