package connections

import (
	"bufio"
	"log"
	"net"

	"github.com/XDBerry29/monitor-app/services"
	"github.com/XDBerry29/monitor-app/utils"
)

type pipeProccesConnection struct {
	name       string
	conn       net.Conn
	logService *services.LogService
	sendFlag   bool
}

func NewPipeConnection(name string, conn net.Conn, logService *services.LogService) ProccesConnection {
	return &pipeProccesConnection{
		name:       name,
		conn:       conn,
		logService: logService,
		sendFlag:   true,
	}
}

// GetName implements ProccesConnection.
func (p *pipeProccesConnection) GetName() string {
	return p.name
}

// Listen implements ProccesConnection.
func (p *pipeProccesConnection) Listen() error {
	defer p.conn.Close()
	scanner := bufio.NewScanner(p.conn)

	for scanner.Scan() {
		message := scanner.Bytes()
		pLog, err := utils.CreateLog(message)
		if err != nil {
			log.Printf("Failed to create log: %v", err)
			return err
		}
		err = p.logService.ProccesLog(pLog, p.sendFlag)
		if err != nil {
			log.Printf("Failed to process log: %v", err)
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		// Check for EOF error to detect connection closure
		if err == bufio.ErrFinalToken {
			log.Println("Connection closed by client.")
		} else {
			log.Printf("Scanner error: %v", err)
		}
		return err
	}

	log.Println("Connection closed normally.")
	return nil
}

// SwitchTransmiFlag implements ProccesConnection.
func (p *pipeProccesConnection) SwitchTransmiFlag() {
	p.sendFlag = !p.sendFlag
}

func (p *pipeProccesConnection) GetSendFlag() bool {
	return p.sendFlag
}
