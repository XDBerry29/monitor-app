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
	wsService := services.NewWsService()
	logservice := services.NewLogService(logRepo, wsService)
	hub := connections.NewConnectionHub()
	pipeName := os.Getenv("PIPE_NAME")
	pipeServer := connections.NewPipeConnectionReciver(pipeName, hub, logservice)
	go logservice.UpdateFileOnNewDay(DIR)
	go pipeServer.ListenNewConnection()

	wsController := controller.NewWsController(wsService)
	routes.InitWsRoutes(wsController, server)

	PORT := os.Getenv("PORT")
	server.Logger.Fatal(server.Start(":" + PORT))
}
