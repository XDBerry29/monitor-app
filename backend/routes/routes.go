package routes

import (
	"github.com/XDBerry29/monitor-app/controller"
	"github.com/labstack/echo/v4"
)

func InitWsRoutes(c *controller.WsController, e *echo.Echo) {
	e.GET("/ws", c.HandleWs)
}
