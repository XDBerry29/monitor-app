package connections

import (
	"fmt"
)

type connectionHub struct {
	hub map[string]ProccesConnection
}

func NewConnectionHub() ConnectionHub {
	return &connectionHub{
		hub: make(map[string]ProccesConnection),
	}
}

// AddConnection implements ConnectionHub.
func (p *connectionHub) AddConnection(connection ProccesConnection) error {
	p.hub[connection.GetName()] = connection
	return nil
}

// DeleteConnection implements ConnectionHub.
func (p *connectionHub) DeleteConnection(connection ProccesConnection) error {
	delete(p.hub, connection.GetName())
	return nil
}

// GetConectionByName implements ConnectionHub.
func (p *connectionHub) GetConectionByName(connectionName string) (ProccesConnection, error) {
	resultConnection := p.hub[connectionName]
	if resultConnection == nil {
		return nil, fmt.Errorf("connection not found")
	}

	return resultConnection, nil
}

// GetAllConnections implements ConnectionHub.
func (p *connectionHub) GetAllConnections() map[string]ProccesConnection {
	return p.hub
}
