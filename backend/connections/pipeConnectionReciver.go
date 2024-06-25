package connections

import (
	"bufio"
	"log"
	"net"

	"github.com/XDBerry29/monitor-app/services"
	"github.com/XDBerry29/monitor-app/utils"
	"github.com/natefinch/npipe"
)

type pipeConnectionReciver struct {
	pipeName   string
	hub        ConnectionHub
	logService *services.LogService
}

func NewPipeConnectionReciver(pipename string, hub ConnectionHub, logService *services.LogService) ConnectionReciver {
	return &pipeConnectionReciver{
		pipeName:   pipename,
		hub:        hub,
		logService: logService,
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
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	conn_message, err := utils.CreateConnectionMessage(message)
	if err != nil {
		return err
	}

	pipeConn := NewPipeConnection(conn_message.ProcessName, conn, p.logService)
	p.hub.AddConnection(pipeConn)
	defer p.hub.DeleteConnection(pipeConn)

	 pipeConn.Listen()

	return nil

}
