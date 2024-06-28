package services

import (
	"github.com/XDBerry29/monitor-app/models"
	"github.com/gorilla/websocket"
)

type Message interface {
	models.ConnectionMessage | models.Log
}

type WsService[M Message] interface {
	AddWs(id string, ws *websocket.Conn)
	RemoveWs(id string)
	SendAll(message M)
	SendOne(message M, id string)
	Listen(id string)
}
