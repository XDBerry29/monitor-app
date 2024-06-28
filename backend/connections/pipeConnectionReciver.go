package connections

import (
	"bufio"
	"log"
	"net"
	"time"

	"github.com/XDBerry29/monitor-app/services"
	"github.com/XDBerry29/monitor-app/utils"
	"github.com/natefinch/npipe"
)

type pipeConnectionReciver struct {
	pipeName      string
	hub           ConnectionHub
	logService    *services.LogService
	connService   *services.ConnectionService
	clientService *services.ClientService
}

func NewPipeConnectionReciver(pipename string, hub ConnectionHub, logService *services.LogService, connService *services.ConnectionService, clientService *services.ClientService) ConnectionReciver {
	return &pipeConnectionReciver{
		pipeName:      pipename,
		hub:           hub,
		logService:    logService,
		connService:   connService,
		clientService: clientService,
	}
}

// ListenNewConnection implements ConnectionReciver.
func (p *pipeConnectionReciver) ListenNewConnection() error {
	ln, err := npipe.Listen(p.pipeName)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go p.HandleConnection(conn)
	}

}

// HandleConnection implements ConnectionReciver.
func (p *pipeConnectionReciver) HandleConnection(conn net.Conn) error {
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		return scanner.Err()
	}

	initialMessage := scanner.Bytes()

	conn_message, err := utils.CreateConnectionMessageNewConn(initialMessage)
	if err != nil {
		return err
	}

	pipeConn := NewPipeConnection(conn_message.ProcessName, conn, p.logService)

	p.hub.AddConnection(pipeConn)
	p.connService.ProccesConnectionMessage(conn_message)
	defer p.hub.DeleteConnection(pipeConn)

	p.clientService.AddOnFilterToAllClients(conn_message.ProcessName)
	defer p.clientService.AddOnFilterToAllClients(conn_message.ProcessName)

	pipeConn.Listen()
	now := time.Now()
	formatTime := now.Format("15:04:05")
	msg := utils.CreateConnectionMessage(pipeConn.GetName(), formatTime, false)
	p.connService.ProccesConnectionMessage(msg)

	return nil

}
