package controller

import (
	"net/http"

	"github.com/XDBerry29/monitor-app/services"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WsController struct {
	wsService *services.WsService
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWsController(wsService *services.WsService) *WsController {
	return &WsController{wsService: wsService}
}

func (c *WsController) HandleWs(context echo.Context) error {
	ws, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		return err
	}

	c.wsService.AddWs(ws)

	return nil
}
