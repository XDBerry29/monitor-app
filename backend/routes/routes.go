package routes

import (
	"github.com/XDBerry29/monitor-app/controller"
	"github.com/labstack/echo/v4"
)

func InitWsRoutes(wsC *controller.WsController, conC *controller.ConnectionController, e *echo.Echo) {
	e.GET("/ws", wsC.HandleWs)
	e.GET("/wsp", conC.HandleWs)
}
