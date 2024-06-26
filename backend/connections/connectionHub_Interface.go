package connections

type ConnectionHub interface {
	AddConnection(connection ProccesConnection) error
	DeleteConnection(connection ProccesConnection) error
	GetConectionByName(connectionName string) (ProccesConnection, error)
	GetAllConnections() map[string]ProccesConnection
}
