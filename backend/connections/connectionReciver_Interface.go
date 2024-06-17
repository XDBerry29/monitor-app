package connections

import "net"

type ConnectionReciver interface {
	ListenNewConnection() error
	HandleConnection(conn net.Conn) error
}
