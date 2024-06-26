package controller

import (
	"encoding/json"

	"github.com/XDBerry29/monitor-app/connections"
	"github.com/XDBerry29/monitor-app/services"
	"github.com/XDBerry29/monitor-app/utils"
	"github.com/labstack/echo/v4"
)

type ConnectionController struct {
	wsService *services.WsService
	conHub    connections.ConnectionHub
}

func NewConnectionController(wsService *services.WsService, conHub connections.ConnectionHub) *ConnectionController {
	return &ConnectionController{wsService: wsService, conHub: conHub}
}

// HandleWs handles websocket connections
func (c *ConnectionController) HandleWs(context echo.Context) error {
	ws, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		return err
	}

	proccesConnection := c.conHub.GetAllConnections()

	for _, connection := range proccesConnection {
		message := utils.CreateConnectionMessage(connection.GetName(), "", true, connection.GetSendFlag())
		sMessage, err := json.Marshal(message)
		if err != nil {
			return err
		}
		c.wsService.SendOne(sMessage, ws)
	}

	c.wsService.AddWs(ws)

	return nil

}
