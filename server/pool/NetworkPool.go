package pool

import "net"

type ClientPool struct {
	Conn   net.Conn
	ConnID string
}

const (
	MAX_CONN = 2000
)

type NetPool map[net.Conn]ClientPool

var clients = make(NetPool, 0)

func AddClient(conn net.Conn) {
	clients[conn] = ClientPool{
		Conn:   conn,
		ConnID: "dsjlksdhds",
	}
}

func DisconnectClient(conn net.Conn) {
	delete(clients, conn)
}

func GetClients() NetPool {
	return clients
}

func IsPoolMaxed() bool {
	return len(clients) == MAX_CONN
}
