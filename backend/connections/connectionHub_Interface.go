package connections

type ConnectionHub interface {
	AddConnection(connection ProccesConnection) error
	DeleteConnection(connection ProccesConnection) error
	GetConectionByName(connectionName string) (ProccesConnection, error)
}
