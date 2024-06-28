package routes

import (
	"github.com/XDBerry29/monitor-app/controller"
	"github.com/labstack/echo/v4"
)

func InitWsRoutes(logC *controller.LogController, conC *controller.ConnectionController, e *echo.Echo) {
	e.GET("/ws", logC.HandleWs)
	e.GET("/wsp", conC.HandleWs)
}

func InitApiRoutes(conC *controller.ConnectionController, e *echo.Echo) {
	// api := e.Group("/api")
	// api.PUT("/connfilter",)
	// api.PUT("/logfilter",)
	// api.GET("/logfilter",)
}
