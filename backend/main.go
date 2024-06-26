package main

import (
	"log"
	"os"

	"github.com/XDBerry29/monitor-app/connections"
	"github.com/XDBerry29/monitor-app/controller"
	"github.com/XDBerry29/monitor-app/repsitories"
	"github.com/XDBerry29/monitor-app/routes"
	"github.com/XDBerry29/monitor-app/services"
	"github.com/XDBerry29/monitor-app/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading the env...")
	}
	server := echo.New()
	server.Use(middleware.Recover())
	server.Use(middleware.Logger())

	//init the log server
	DIR := os.Getenv("LOG_FILE_DIR")
	file, err := utils.CreateLogFile(DIR)
	if err != nil {
		log.Fatal("Error creating file")
	}

	logRepo := repsitories.NewLogRepoFile(file)

	wsServiceLogs := services.NewWsService()
	wsServiceConn := services.NewWsService()

	logservice := services.NewLogService(logRepo, wsServiceLogs)
	connService := services.NewConnectionService(wsServiceConn)

	hub := connections.NewConnectionHub()
	pipeName := os.Getenv("PIPE_NAME")
	pipeServer := connections.NewPipeConnectionReciver(pipeName, hub, logservice, connService)
	go logservice.UpdateFileOnNewDay(DIR)
	go pipeServer.ListenNewConnection()

	connController := controller.NewConnectionController(wsServiceConn, hub)
	wsController := controller.NewWsController(wsServiceLogs)
	routes.InitWsRoutes(wsController, connController, server)

	PORT := os.Getenv("PORT")
	server.Logger.Fatal(server.Start(":" + PORT))
}
